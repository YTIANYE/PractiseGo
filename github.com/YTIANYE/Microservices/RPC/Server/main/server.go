package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

// 定义类对象
type World struct {
}

// 绑定类方法
func (this *World) HelloWorld(name string, resp *string) error {
	*resp = name + " 你好！"
	return nil // 返回值不等于nil出错，等于nil不出错
	// return errors.New("未知错误！") // 假如有错误,存在返回的值也不会显示
}

func main() {
	// 1.注册RPC服务，绑定对象方法
/*	err := rpc.RegisterName("hello", new(World))
	if err != nil {
		fmt.Println("注册RPC服务失败！", err)
		return
	}*/

	// 封装后的注册
	RegisterService(new(World))

	// 2.设置监听
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("new.Listen err", err)
		return
	}
	defer listener.Close()

	fmt.Println("开始监听 ...")
	// 3.建立连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Accept() err:", err)
		return
	}
	defer conn.Close()
	fmt.Println("连接建立成功 ...")
	// 4.绑定服务
	// rpc.ServeConn(conn)
	jsonrpc.ServeConn(conn)
}
