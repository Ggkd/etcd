package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 3,
	})
	if err != nil {
		fmt.Println("new client err :", err)
		return
	}
	fmt.Println("conn etcd success")

	lease, err := client.Grant(context.Background(), 5)
	if err != nil {
		fmt.Println("grant lease err:", err)
		return
	}

	_, err = client.Put(context.Background(), "age", "11", clientv3.WithLease(lease.ID))
	if err != nil {
		fmt.Println("put err:", err)
		return
	}
	fmt.Println("put success")
}
