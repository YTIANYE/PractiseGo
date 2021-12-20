package main

import "fmt"

/*func main() {
	// 1. 用RPC连接服务器  --Dial
	// conn, err := rpc.Dial("tcp", "127.0.0.1:8800")
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer conn.Close()

	// 2.调用远程函数
	var reply string // 接收返回值---传出参数
	err = conn.Call("hello.HelloWorld", "张三", &reply)
	if err != nil{
		fmt.Println("Call:", err)
		return
	}
	fmt.Println(reply)
}*/

// 结合 design.go 去测试
func main(){
	myClient := InitClient("127.0.0.1:8800")
	var resp string
	err := myClient.HelloWorld("李四", &resp)
	if err != nil{
		fmt.Println("HelloWorld err:", err)
		return
	}
	fmt.Println(resp, err)

}
