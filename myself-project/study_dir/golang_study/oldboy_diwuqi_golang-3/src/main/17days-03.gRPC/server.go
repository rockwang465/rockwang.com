package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	pb "oldboy_diwuqi_golang-3/src/main/17days-03.gRPC/proto"
)

// 定义空接口
type UserInfoService struct{}

var u = UserInfoService{} // 初始化

// 实现方法
// (ctx context.Context, req *pb.UserRequest)(*pb.UserResponse, error)来自user.pb.go中297行的GetUserInfo方法中的参数及返回值
func (s *UserInfoService) GetUserInfo(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	// 通过用户名查询用户信息
	name := req.Name
	// 数据里查用户信息
	if name == "zs" {
		resp = &pb.UserResponse{
			Id:    1,
			Name:  name,
			Age:   22,
			Hobby: []string{"Sing", "Run"},
		}
	}
	return

}

func main() {
	// 地址
	addr := "127.0.0.1:8080"
	// 1.需要监听
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Error: net listen err = ", err)
	}
	fmt.Println("Info : listen is successfully")

	// 2.需要实例化gRPC服务端
	s := grpc.NewServer()

	// 3.在gRPC上注册微服务
	// 在user.pb.go的301行有注册函数: RegisterUserInfoServiceServer(s *grpc.Server, srv UserInfoServiceServer)
	pb.RegisterUserInfoServiceServer(s, &u)

	// 4.启动微服务端
	s.Serve(listener)
}
