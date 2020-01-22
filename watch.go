package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:[]string{"127.0.0.1:2379"},
		DialTimeout:time.Second*3,
	})
	if err != nil {
		fmt.Println("new client err :", err)
		return
	}
	fmt.Println("conn etcd success")

	watchChan := client.Watch(context.Background(), "name")
	for event := range watchChan {
		for _, ev := range event.Events {
			fmt.Printf("type:%s,  key:%s, value:%s \n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}