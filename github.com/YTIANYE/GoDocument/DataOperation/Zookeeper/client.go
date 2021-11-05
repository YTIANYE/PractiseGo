package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)


func clientcheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	for i := 0; i < 100; i++ {
		startClient()

		time.Sleep(1 * time.Second)
	}
}

func startClient() {
	// service := "127.0.0.1:8899"
	//获取地址
	serverHost, err := getServerHost()
	if err != nil {
		fmt.Printf("get server host fail: %s \n", err)
		return
	}

	fmt.Println("connect host: " + serverHost)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", serverHost)
	clientcheckError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	clientcheckError(err)
	defer conn.Close()

	// 获取server发送的时间后断开
	_, err = conn.Write([]byte("timestamp"))
	clientcheckError(err)

	result, err := ioutil.ReadAll(conn)
	clientcheckError(err)
	fmt.Println(string(result))

	return
}

func getServerHost() (host string, err error) {
	conn, err := clientGetConnect()
	if err != nil {
		fmt.Printf(" connect zk error: %s \n ", err)
		return
	}
	defer conn.Close()

	serverList, err := GetServerList(conn)
	if err != nil {
		fmt.Printf(" get server list error: %s \n", err)
		return
	}

	count := len(serverList)
	if count == 0 {
		err = errors.New("server list is empty \n")
		return
	}

	//随机选中一个返回
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	host = serverList[r.Intn(3)]
	return
}

func clientGetConnect() (conn *zk.Conn, err error) {
	zkList := []string{"localhost:2181"}
	conn, _, err = zk.Connect(zkList, 10*time.Second)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetServerList(conn *zk.Conn) (list []string, err error) {
	list, _, err = conn.Children("/go_servers")
	return
}