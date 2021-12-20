package main

import (
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 客户端的封装

// 像调用本地函数一样，调用远程函数

// 定义一个类
type Myclient struct {
	c *rpc.Client
}

// 由于使用了c去调用Call ,因此需要初始化c
func InitClient(addr string) Myclient {
	conn, _ := jsonrpc.Dial("tcp", addr)
	return Myclient{c: conn}
}

// 实现函数， 原型参照Interface来实现

func (this *Myclient) HelloWorld(a string, b *string) error {
	// 参数1 参照Interface RegisterName 而来
	// a 传入参数  b 传出参数
	return this.c.Call("hello.HelloWorld", a, b)
}
