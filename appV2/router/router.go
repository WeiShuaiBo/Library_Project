package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/appV2/aliPay"
	//swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/files"
	//ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger"
	"io"
	_ "library/appV2/docs"
	"library/appV2/logic"
	"library/appV2/model"
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
	r.GET("/html", logic.Html)
	r.GET("/:pageNumber", logic.GetAll)
	r.GET("/qrcode", logic.Qrcode)
	r.GET("/AutomaticLogin/:code", logic.AutomaticLogin)

	//r.GET("/GetCode", logic.SendNum)            //
	r.POST("/Captcha", logic.Captcha)
	r.POST("/Registered/:code", logic.Registered)

	//支付回调
	r.GET("/alipay/callback", func(c *gin.Context) {
		aliPay.Callback(c.Writer, c.Request)
	})
	r.GET("/alipay/notify", func(c *gin.Context) {
		aliPay.Notify(c.Writer, c.Request)
	})
	r.GET("/aLiPay/:ordersId", logic.ALiPay)

	r.POST("/userLogin", logic.UserLogin) //
	r.GET("/userLogout", logic.Logout)    //
	//r.POST("/users", logic.AddUser)             //
	r.POST("/adminLogin", logic.LibrarianLogin) //
	r.GET("/adminLogout", logic.AdminLogout)    //
	////游客可以浏览书籍和分类
	//r.GET("/books", logic.SearchBook)          //
	r.GET("/books/:id", logic.GetBook)                 //
	r.GET("/book/:bookName", logic.GetBookPhotoByName) //

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 启动定时器
	go func() {
		//固定的，注册swagger api用的
		timer := time.NewTicker(time.Second * 5)

		// 在定时器的通道中接收定时触发的事件
		for range timer.C {
			ret := model.GetBuyBook()

			fmt.Println("定时器触发")
			for _, ret := range ret {
				fmt.Println("遍历订单")
				// 这里可以执行你想要的操作
				createdAt, err := time.Parse(time.RFC3339, ret.CreatedAt)
				if err != nil {
					// 处理解析错误
				}

				if ret.OrderStatus == "待支付" && time.Now().Sub(createdAt) > time.Second*3600 {
					// 执行操作
					if !model.NoticeOrders(ret.OrderId) {
						fmt.Println("noticeOrders出现问题")
					}
				} else {
					if ret.OrderStatus == "未通知尽快支付" || ret.OrderStatus == "已通知尽快支付" {
						model.CancelBuyBook(ret.OrderId)

					}
				}

			}
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
