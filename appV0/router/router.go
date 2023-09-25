package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "library/appV0/docs"
	"library/appV0/logic"
	"library/appV0/model"
	//"time"
)

func New() *gin.Engine {
	// cd .\LM_V1\
	// http://localhost:8083/swagger/index.html
	model.New()
	r := gin.Default()
	r.LoadHTMLGlob("E:/workspase/go/golandWorkspace/library/appV0/static/*.html")
	r.Static("/static", "E:/workspase/go/golandWorkspace/library/appV0/static")

	userRouter(r)
	adminRouter(r)
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	////验证码
	r.GET("/html", logic.Html)
	r.GET("/:pageNumber", logic.GetAll)

	//r.GET("/GetCode", logic.SendNum)            //
	r.POST("/Captcha", logic.Captcha)
	r.POST("/Registered/:code", logic.Registered)

	r.POST("/userLogin", logic.UserLogin) //
	r.GET("/userLogout", logic.Logout)    //
	//r.POST("/users", logic.AddUser)             //
	r.POST("/adminLogin", logic.LibrarianLogin) //
	r.GET("/adminLogout", logic.AdminLogout)    //
	////游客可以浏览书籍和分类
	//r.GET("/books", logic.SearchBook)          //
	r.GET("/books/:id", logic.GetBook)                 //
	r.GET("/book/:bookName", logic.GetBookPhotoByName) //
	//r.GET("/categories", logic.SearchCategory) //
	/*time1 := time.NewTicker(1 * time.Minute)
	go func() {
		for {
			select {
			case t := <-time1.C:
				CheckRecord() //
				fmt.Printf("定时器1正在运行中...%+v\n", t.Unix())
			}
		}
	}()

	go func() {
		for {
			select {
			case t := <-time1.C:
				WillReturn() //
				fmt.Printf("定时器2正在运行中...%+v\n", t.Unix())
			}
		}
	}()*/
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//固定的，注册swagger api用的
	return r
}

//func CheckRecord() {
//	now := time.Now()
//	fmt.Printf("现在的时间点：%+v\n", now)
//	next := now.Add(5 * time.Minute)
//	fmt.Printf("打印next:%+v\n", next)
//
//	//
//	ret := model.ReturnTime(now)
//	fmt.Println(ret)
//	//
//	ans := model.BanUsers(ret)
//	fmt.Printf("打印ans%+v\n", ans)
//}
//
//func WillReturn() {
//	now := time.Now()
//
//	next := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
//
//	// 计算当前时间与下一个零点的时间差
//	duration := next.Sub(now)
//	time.Sleep(duration)
//	//
//	model.AdvanceOneDay()
//
//}
