package main

import (
	pb "awesomeProject1/Microservices/PB"
	"context" // 上下文，---goroutine (go程) 之间用来进行数据传递的API包
	"fmt"
	"google.golang.org/grpc"
	"net"
)

// 定义类  输入对象
type Children struct {
}

// 按照接口绑定类方法
func (this *Children) SayHello(ctx context.Context, t *pb.Teacher) (*pb.Teacher, error) {
	t.Name += "is Working!"
	return t, nil
}

func main() {
	// 1. 初始一个 grpc 对象
	grpcServer := grpc.NewServer()

	// 2. 注册服务
	pb.RegisterSayNameServer(grpcServer, new(Children))
	fmt.Println("注册服务成功")

	// 3. 设置监听， 指定 IP、port
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil{
		fmt.Println("Listen err:", err)
		return
	}
	defer listener.Close()

	// 4. 启动服务。---- serve()
	grpcServer.Serve(listener)
}
