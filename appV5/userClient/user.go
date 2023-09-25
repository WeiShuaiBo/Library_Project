package userClient

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"library/appV3/tools"
	pb "library/appV5/proto"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name" form:"name" binding:"required"`
	Pwd  string `json:"pwd" form:"pwd" binding:"required"`
}

// UserLogin godoc
// 接口的名字
//
//	@Summary		用户登录
//	@Description	会执行用户登录操作
//	@Tags			user
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			name	formData	string	true	"用户名"
//	@Param			pwd		formData	string	true	"密码"
//	@response		200,500	{object}	tools.HttpCode{data=Token}
//	@Router			/userLogin [POST]
func UserLogin(c *gin.Context) {
	var user User
	//	绑定并判断
	if c.ShouldBind(&user) != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "绑定失败",
			Data:    struct{}{},
		})
		return
	}
	fmt.Printf("看看user")
	fmt.Print(user)

	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect:%v\n", err)
	}
	defer conn.Close()

	//建立连接
	client := pb.NewUserClient(conn)

	//执行rpc调用(这个方法在服务器端来实现并返回结果）
	resp, _ := client.UserLogin(context.Background(), &pb.UserRequest{Name: "kuangshen", Pwd: "123"}) //发送给服务端
	fmt.Println(resp.GetName())
}

//
//
//func UserLogin(c *gin.Context) {
//	//声明user
//	var user User
//	//	绑定并判断
//	if c.ShouldBind(&user) != nil {
//		c.JSON(http.StatusOK, tools.HttpCode{
//			Code:    tools.UserInfoErr,
//			Message: "绑定失败",
//			Data:    struct{}{},
//		})
//		return
//	}
//	fmt.Printf("看看user")
//	fmt.Print(user)
//	//	查询数据库并判断
//	dbUser := model.GetUser(user.Name, user.Pwd)
//	fmt.Println(dbUser.Id)
//	if dbUser.Id <= 0 {
//		c.JSON(http.StatusOK, tools.HttpCode{
//			Code:    tools.UserInfoErr,
//			Message: "查询数据库失败",
//			Data:    struct{}{},
//		})
//		return
//	}
//	//	下发Token并判断
//	a, r, err := tools.Token.GetToken(dbUser.Id, dbUser.Name, "user")
//	if err != nil {
//		c.JSON(http.StatusOK, tools.HttpCode{
//			Code:    tools.UserInfoErr,
//			Message: "生成Token失败",
//			Data:    struct{}{},
//		})
//		return
//	}
//	//	通知成功
//	c.JSON(http.StatusOK, tools.HttpCode{
//		Code:    tools.OK,
//		Message: "登录成功，正在跳转",
//		Data: Token{
//			AccessToken:  a,
//			RefreshToken: r,
//		},
//	})
//	return
//}
