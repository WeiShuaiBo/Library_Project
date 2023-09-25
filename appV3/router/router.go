package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"io"
	"library/appV3/aliPay"
	_ "library/appV3/docs"
	"library/appV3/logic"
	"library/appV3/model"
	"net/http"
	"os"
	"sync"
	"time"
	//"time"
)

func New() *gin.Engine {
	// cd .\LM_V1\
	// http://localhost:8083/swagger/index.html
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	gin.ForceConsoleColor()

	//以上都是日志部分
	model.New()
	r := gin.Default()

	// 添加跨域中间件
	//r.Use(cors.Default())

	limiter := &RateLimiter{
		rateLock: sync.Mutex{},
		rateMap:  make(map[string]int),
	}

	x := 5  // 时间窗口为 x 秒
	y := 30 // 最大请求数量为 y
	//z := 30 // 禁止访问时间为 z 秒
	r.Use(limiter.Middleware(y, time.Second*time.Duration(x)))

	r.LoadHTMLGlob("E:/workspase/go/golandWorkspace/library/appV0/static/*.html")
	r.Static("/static", "E:/workspase/go/golandWorkspace/library/appV0/static")

	r.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	userRouter(r)
	adminRouter(r)
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	////验证码
	r.GET("/html", logic.Html)                           //打开前端页面
	r.GET("/:pageNumber", logic.GetAll)                  //分页查询全部图书
	r.GET("/qrcode", logic.Qrcode)                       //获取二维码完成单点登录
	r.GET("/AutomaticLogin/:code", logic.AutomaticLogin) //获取二维码后自动请求结果

	//r.GET("/GetCode", logic.SendNum)            //
	r.POST("/Captcha", logic.Captcha)             //注册前请求电子邮箱验证码或者手机验证码
	r.POST("/Registered/:code", logic.Registered) //注册

	//支付回调
	r.GET("/alipay/callback", func(c *gin.Context) {
		aliPay.Callback(c.Writer, c.Request)
	}) //支付成功阿里支付宝回调
	r.GET("/alipay/notify", func(c *gin.Context) {
		aliPay.Notify(c.Writer, c.Request)
	})                                       //支付失败的回调
	r.GET("/aLiPay/:ordersId", logic.ALiPay) //支付时用户调的接口

	r.POST("/userLogin", logic.UserLogin) //用户登录
	r.GET("/userLogout", logic.Logout)    //取消登录
	//r.POST("/users", logic.AddUser)             //
	r.POST("/adminLogin", logic.LibrarianLogin) //管理员登录
	r.GET("/adminLogout", logic.AdminLogout)    //管理园退出登录
	////游客可以浏览书籍和分类
	//r.GET("/books", logic.SearchBook)          //
	r.GET("/books/:id", logic.GetBook)                 //根据id得到图书
	r.GET("/book/:bookName", logic.GetBookPhotoByName) //根据书名模糊查询图书获得图书封面

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) //swagger页面

	client := model.CreateRedisClient()
	// 订阅指定的通道
	ctx := context.Background()
	pubsub := client.Subscribe(ctx, "channel-name")
	//pubsub := client.Subscribe(ctx, "your-queue")
	_, err := pubsub.Receive(ctx)
	if err != nil {
		fmt.Printf("订阅通道失败：%v\n", err)
		return r
	}

	// 启动后台 goroutine 处理订阅消息
	go func() {
		ch := pubsub.Channel()
		for msg := range ch {
			// 处理接收到的消息
			fmt.Println("接收到的消息:", msg.Payload)

		}
	}()
	return r
}

// NewRequestLimiter函数用于创建一个新的请求频率限制器，并返回一个指向该结构体的指针。
type RateLimiter struct {
	rateLock sync.Mutex
	rateMap  map[string]int
}

func (rl *RateLimiter) Middleware(maxRequests int, interval time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()         //获取ip的方法
		rl.rateLock.Lock()         //互斥锁上锁
		defer rl.rateLock.Unlock() //完成后互斥锁解锁

		count, exists := rl.rateMap[ip] //count为ip的计数，exists为是否存在，存在为true
		if !exists {
			rl.rateMap[ip] = 1
		} else {
			rl.rateMap[ip]++
		}
		if count > maxRequests {
			c.JSON(429, gin.H{"error": "请求频率过高，请稍后再试"})
			c.Abort()
		}

		go func() {
			time.Sleep(interval)
			rl.rateLock.Lock()
			defer rl.rateLock.Unlock()
			rl.rateMap[ip]--
		}()

		c.Next()
	}
}
