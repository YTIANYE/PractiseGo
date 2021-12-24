package main

import (
	"awesomeProject1/Microservices/consul/pb"
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"strconv"
)

func main(){

	// 1. 初始化 consul 配置
	consulConfig := api.DefaultConfig()

	// 2. 创建 consul对象  --(可以重新指定consul属性，ip,port,...,也可以使用默认)
	consulClient, err := api.NewClient(consulConfig)

	// 3. 服务发现， 从consul上获取健康的服务
	services, _, err := consulClient.Health().Service("grpc and consul", "grpc", true, nil )

	// 这里可以添加简单的负载均衡，访问压力均摊给集群中的每个服务

	addr := services[0].Service.Address + ":" + strconv.Itoa(services[0].Service.Port)

	// grpc 远程调用

	// 1. 连接服务
	// 原始的方式
	// grpcConn, _ := grpc.Dial("127.0.0.1:8800", grpc.WithInsecure())
	// 使用服务发现consul上的 ip和port 来与服务建立连接
	grpcConn, _ := grpc.Dial(addr, grpc.WithInsecure())

	// 2. 初始化一个grpc客户端
	grpcClient := pb.NewHelloClient(grpcConn)

	var person pb.Person
	person.Name = "张三"
	person.Age = 18

	// 3. 调用远程函数
	p, err := grpcClient.SayHello(context.TODO(), &person)
	fmt.Println(p, err)
}
