package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/appV3/logic"
	"library/appV3/model"
	"library/appV3/tools"
	"net/http"
	"strconv"
)

func adminRouter(r *gin.Engine) {
	//librarian := r.Group("/librarians").Use(librarianCheck())
	//      /admin/users
	base := r.Group("/admin")

	//r.StaticFS("/static", http.Dir("static"))
	base.Use(librarianCheck())
	//user := base.Group("/users")
	//{
	//user.GET("/:id", logic.GetUser)
	//user.GET("", logic.SearchUser)            //
	//user.PUT("/:id", logic.UpdateUserByAdmin) //
	//user.DELETE("/:id", logic.DeleteUser)     //
	//获取该用户已归还或者未归还的所有记录
	//user.GET("/:id/records/:status", logic.GetUserBook) //
	//user.POST("/:id/records/:bookId", logic.BorrowBook)
	//user.PUT("/:id/records/:bookId", logic.ReturnBook)
	//}
	//书的所有资源
	//    /admin/books
	book := base.Group("/books")
	//{
	//	book.GET("/:id", logic.GetBook) // 直接使用谁都可以查看图书，此路径先不用
	book.POST("/AddBook", logic.AddBook)
	book.GET("/DeleteBook/:bookId", logic.DeleteBook)
	//	//book.GET("", logic.SearchBook)
	//	book.POST("", logic.AddBook)          //
	//	book.PUT("/:id", logic.UpdateBook)    //
	//	book.DELETE("/:id", logic.DeleteBook) //
	//}

	//   /admin/categories
	//category := base.Group("/categories")
	//{
	//	category.GET("/:id", logic.GetCategory) //这个不必要写
	//	//category.GET("", logic.SearchCategory)
	//	category.POST("", logic.AddCategory)          //
	//	category.PUT("/:id", logic.UpdateCategory)    //
	//	category.DELETE("/:id", logic.DeleteCategory) //
	//}
	//记录表的资源  /admin/records
	record := base.Group("/records")
	{
		//所有借书还书记录
		record.GET("/getRecords", logic.GetRecords) //
		//所有归还或者未归还的记录
		record.GET("/status/:id", logic.GetUserRecordStatus) //
		record.GET("/GetPersonalInformation/:userId", logic.GetUserInformation)
		record.GET("/GETBookRecord/:bookId", logic.GETBookRecord)
	}
}
func librarianCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Print("检查session中......")
		CookieName, _ := c.Cookie("name")
		CookieId, _ := c.Cookie("id")
		SessionID, _ := c.Cookie("sessionID")
		data, err1 := model.GetSession(SessionID)
		if err1 != nil {
			fmt.Println("getSession失败")
		}
		fmt.Print(data)
		id := data["user_id"]
		name := data["user_name"]

		idInt64, idErr := strconv.ParseInt(fmt.Sprintf("%v", id), 10, 64)
		if idErr != nil {
			fmt.Println("ID转换失败：", idErr)
			c.AbortWithStatusJSON(http.StatusInternalServerError, tools.HttpCode{
				Code:    tools.OK,
				Message: "Sid错误",
			})
			return
		}
		fmt.Printf("cookieName：%d，类型：%T\n", CookieName, CookieName)
		fmt.Printf("CookieId：%d，类型：%T\n", CookieId, CookieId)
		fmt.Printf("SessionName：%d，类型：%T\n", name, name)
		fmt.Printf("SessionId：%d，类型：%T\n", id, id)
		CId, ciErr := strconv.ParseInt(CookieId, 10, 64)
		if ciErr != nil {
			fmt.Println("CookieID转换失败：", ciErr)
			c.AbortWithStatusJSON(http.StatusUnauthorized, tools.HttpCode{
				Code:    tools.NotLogin,
				Message: "用户信息获取失败",
			})
			return
		}

		//Cname, cnErr := strconv.ParseInt(CookieName, 10, 64)
		//if cnErr != nil {
		//	fmt.Println("Cookie名称转换失败：", cnErr)
		//	c.AbortWithStatusJSON(http.StatusUnauthorized, tools.HttpCode{
		//		Code:    tools.NotLogin,
		//		Message: "用户信息获取失败",
		//	})
		//	return
		//}

		//Sname, snOk := name.(string)
		//if !snOk {
		//	fmt.Println("Session名称转换失败")
		//	c.AbortWithStatusJSON(http.StatusInternalServerError, tools.HttpCode{
		//		Code:    tools.OK,
		//		Message: "服务器Sname内部错误",
		//	})
		//	return
		//}

		//fmt.Printf("cookieName：%d，类型：%T\n", Cname, Cname)
		fmt.Printf("CookieId：%d，类型：%T\n", CId, CId)
		//fmt.Printf("name：%s，类型：%T\n", Sname, Sname)
		fmt.Printf("id：%d，类型：%T\n", idInt64, idInt64)

		//if !ok1 || !ok2 || idInt64 <= 0 || Sname != CookieName || CId != idInt64 {
		//	c.AbortWithStatusJSON(http.StatusUnauthorized, tools.HttpCode{
		//		Code:    tools.NotLogin,
		//		Message: "用户信息获取失败",
		//	})
		//	return
		//}
		fmt.Printf("session通过")
		c.Set("name", name)
		c.Set("id", id)
		c.Next()
	}
}

//func librarianCheck() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		fmt.Print("检查session中......")
//		CookieName, _ := c.Cookie("name")
//		CookieId, _ := c.Cookie("id")
//		data := model.GetSession(c)
//		id, ok1 := data["id"]
//		name, ok2 := data["name"]
//
//		idInt64, _ := id.(int64)
//		CId, _ := strconv.ParseInt(CookieId, 10, 64)
//		//CId, _ := CookieId.(int64)
//		Cname, _ := strconv.ParseInt(CookieName, 10, 64)
//		//Cname, _ := CookieName.(int64)
//		Sname, _ := name.(int64)
//
//		fmt.Printf("cookieName：%s，类型： %T\n", Cname)
//		fmt.Printf("CookieId：%d，类型： %T\n", CId)
//		fmt.Printf("name：%d，类型： %T\n", Sname)
//		fmt.Printf("id：%d，类型： %T\n", id)
//
//		if !ok1 || !ok2 || idInt64 <= 0 || Sname != Cname || CId != id {
//			c.AbortWithStatusJSON(http.StatusUnauthorized, tools.HttpCode{
//				Code:    tools.NotLogin,
//				Message: "用户信息获取失败",
//			})
//			return
//		}
//		c.Next()
//
//		c.Set("name", name)
//		c.Set("id", idInt64)
//		c.Next()
//	}
//}
