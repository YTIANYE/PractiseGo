package main

import (
	test66 "awesomeProject1/Microservices/gin/proto/test66"
	context2 "context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
)

func CallRemote(context *gin.Context) {
	// 1 初始化客户端
	microClient := test66.NewTest66Service("go.micro.srv.test66", client.DefaultClient)
	fmt.Println()

	// 2 调用远程服务
	resp, err := microClient.Call(context2.TODO(), &test66.Request{
		Name:"张三",
	})
	if err != nil{
		fmt.Println("call err:", err)
		return
	}

	// 3 为了方便查看， 在打印之前将结果返回浏览器
	context.Writer.WriteString(resp.Msg)
	fmt.Println(resp, err)

}

func main() {
	// 1 初始化路由/初始化web 引擎
	router := gin.Default()

	// 2 路由匹配
	// router.GET("/", func(context *gin.Context) {
	// 	context.Writer.WriteString("hello world!")
	// })
	router.GET("/", CallRemote)

	// 3 启动运行
	router.Run(":8080")
}
