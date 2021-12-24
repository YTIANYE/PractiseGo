package main

import (
	"awesomeProject1/Microservices/consul/pb"
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"net"
)

// 定义类
type Children struct {
}

// 绑定类方法， 实现接口
func (this *Children) SayHello(ctx context.Context, p *pb.Person) (*pb.Person, error) {
	p.Name = "你好" + p.Name
	return p, nil
}

func main() {

	// 把 grpc服务注册到consul上

	// 1. 初始化consul配置

	consulConfig := api.DefaultConfig()

	// 2. 创建consul对象
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Println("api.NewClient err:", err)
		return
	}

	// 3. 告诉consul即将注册的服务配置信息
	registerService := api.AgentServiceRegistration{
		ID:      "id_1",
		Tags:    []string{"grpc", "consul"}, // 服务别名
		Name:    "grpc and consul",     // 当前服务名字
		Address: "127.0.0.1",
		Port:    8800,
		Check: &api.AgentServiceCheck{
			CheckID:"consul grpc test",
			TCP: "127.0.0.1:8800",
			Timeout:"1s",
			Interval: "5s",
		},
	}

	// 4. 注册服务到consul上
	consulClient.Agent().ServiceRegister(&registerService)


	// grpc 远程调用

	// 1. 初始化 grpc 对象，
	grpcServer := grpc.NewServer()

	// 2.注册服务
	pb.RegisterHelloServer(grpcServer, new(Children))

	// 3.设置监听， 指定IP/port
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("Listen err:", err)
	}
	defer listener.Close()

	// 4.启动服务
	fmt.Println("服务启动...")
	grpcServer.Serve(listener)

}
