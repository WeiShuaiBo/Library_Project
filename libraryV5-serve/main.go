package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "library/libraryV5-serve/proto"
	"net"
)

type server struct {
	pb.UnimplementedUserServer
}

// appid和appkey头的自定义元数据以进行身份验证。
func (s *server) UserLogin(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {

	return &pb.UserResponse{Name: "hello" + req.Name}, nil
}

func main() {

	//开启端口
	listen, _ := net.Listen("tcp", ":9090")
	//创建grpc服务
	grpcServer := grpc.NewServer()
	//在grpc服务端中去注册我们自己编写的服务
	pb.RegisterUserServer(grpcServer, &server{})

	//启动服务
	err := grpcServer.Serve(listen) //这个方法会导致main方法阻塞起来，
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
