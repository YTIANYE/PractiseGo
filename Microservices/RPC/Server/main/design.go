package main

import "net/rpc"

// 要求 服务端在注册RPC对象时，能让编译器检测出 注册对象是否合法
// 避免运行时，发现注册对象不合法

// 创建接口，在接口中定义方法原型
type MyInterface interface {
	HelloWorld(string, *string) error // 限定参数是两个
}

// 封装一个函数实现
// 调用该方法时需要给i传参， 参数应该是实现了HelloWorld方法的类对象！
func RegisterService(i MyInterface) {
	rpc.RegisterName("hello", i)
}
