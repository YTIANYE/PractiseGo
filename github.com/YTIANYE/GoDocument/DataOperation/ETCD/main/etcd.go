package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3/concurrency"

	"github.com/coreos/etcd/clientv3"
	//"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

// etcd client put/get demo
// use etcd/clientv3


/* put */
func etcd_put(cli *clientv3.Client) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := cli.Put(ctx, "lmh", "lmh")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
}

/* get */
func etcd_get(cli *clientv3.Client) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "lmh")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
}

/* watch key:lmh change */
func etcd_watch(cli *clientv3.Client) {

	rch := cli.Watch(context.Background(), "lmh") // <-chan WatchResponse
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}

/* lease租约 */
func etcd_lease(cli *clientv3.Client) {
	// 创建一个5秒的租约
	resp, err := cli.Grant(context.TODO(), 20)
	if err != nil {
		log.Fatal(err)
	}

	// 5秒钟之后, /lmh/ 这个key就会被移除
	_, err = cli.Put(context.TODO(), "lmh", "lmh", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}
	for {
		etcd_get(cli)
	}


}

/*keepAlive*/
func etcd_keep(cli *clientv3.Client) {

	resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}

	// the key 'foo' will be kept forever
	ch, kaerr := cli.KeepAlive(context.TODO(), resp.ID)
	if kaerr != nil {
		log.Fatal(kaerr)
	}
	for {
		ka := <-ch
		fmt.Println("ttl:", ka.TTL)
	}
}

/*基于etcd实现分布式锁*/
func etcd_distributed() {

	cli, err := clientv3.New(clientv3.Config{Endpoints:   []string{"127.0.0.1:2379"}})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// 创建两个单独的会话用来演示锁竞争
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()
	m1 := concurrency.NewMutex(s1, "/my-lock/")

	s2, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s2.Close()
	m2 := concurrency.NewMutex(s2, "/my-lock/")

	// 会话s1获取锁
	if err := m1.Lock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("acquired lock for s1")

	m2Locked := make(chan struct{})
	go func() {
		defer close(m2Locked)
		// 等待直到会话s1释放了/my-lock/的锁
		if err := m2.Lock(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	if err := m1.Unlock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("released lock for s1")

	<-m2Locked
	fmt.Println("acquired lock for s2")
}

func main() {
	/*连接etcd 建立客户端 */
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	//fmt.Printf("cli 的类型 %T", cli)
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)		// 占位符 %v 相应值的默认格式
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()

	// etcd_put(cli)
	// etcd_get(cli)

	// etcd_watch(cli)

	// etcd_lease(cli)
	// etcd_keep(cli)
	etcd_distributed()


}