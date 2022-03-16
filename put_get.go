package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func get(client *clientv3.Client) {
	//get
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := client.Get(ctx, "age")
	if err != nil {
		fmt.Println("get err:", err)
		return
	}
	for _, kv := range result.Kvs {
		fmt.Printf("key: %s,  value:%s \n", kv.Key, kv.Value)
	}
}

func put(client *clientv3.Client) {
	// put
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := client.Put(ctx, "name", "luffy")
	if err != nil {
		fmt.Println("put err:", err)
		return
	}
	fmt.Println("put success")
}

func del(client *clientv3.Client) {
	defer client.Close()
	_, err := client.Delete(context.Background(), "name")
	if err != nil {
		fmt.Println("delete err:", err)
		return
	}
}

func Put2(Client *clientv3.Client) {
	// put
	defer Client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := Client.Put(ctx, "collect_log", `[{"path":"/Users/gongjiabao/Documents/gitdemo/log_collect/test_log/mysql.log","topic":"mysql_log"}]`)
	if err != nil {
		fmt.Printf("put %s err: %v\n", "log_path", err)
		return
	}
	fmt.Println("=====put log_path success=====")
}

func Put1(Client *clientv3.Client) {
	// put
	defer Client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := Client.Put(ctx, "collect_log", `[{"path":"/Users/gongjiabao/Documents/gitdemo/log_collect/test_log/web.log","topic":"web_log"}]`)
	if err != nil {
		fmt.Printf("put %s err: %v\n", "log_path", err)
		return
	}
	fmt.Println("=====put log_path success=====")
}

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

	//put(client)
	get(client)
	//del(client)

	//Put1(client)
	//Put2(client)
}
