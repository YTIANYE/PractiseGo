package main

import (
	pb "awesomeProject1/Microservices/PB"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	// 1. 连接 grpc 服务
	grpcConn, err := grpc.Dial("127.0.0.1:8800", grpc.WithInsecure())
	if err != nil {
		fmt.Println("grpc.Dial err:", err)
		return
	}
	defer grpcConn.Close()

	// 2. 初始化 grpc 客户端
	grpcClient := pb.NewSayNameClient(grpcConn)
	// 创建并初始化一个Teacher对象
	var teacher pb.Teacher
	teacher.Name = "张三"
	teacher.Age = 18
	// 3. 调用远程服务。
	t, err := grpcClient.SayHello(context.TODO(), &teacher)

	fmt.Println(t, err)

}
