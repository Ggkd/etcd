package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
)

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:[]string{"127.0.0.1:2379"},
	})
	if err != nil {
		fmt.Println("new client err :", err)
		return
	}
	fmt.Println("conn etcd success")

	// 创建session来竞争锁
	s1, err := concurrency.NewSession(client)
	if err != nil {
		fmt.Println("new session1 err:", err)
	}
	defer s1.Close()

	s2, err := concurrency.NewSession(client)
	if err != nil {
		fmt.Println("new session2 err:", err)
	}
	defer s2.Close()
	m1 := concurrency.NewMutex(s1, "/lock")
	m2 := concurrency.NewMutex(s2, "/lock")

	// s1获取锁
	if err := m1.Lock(context.Background()); err != nil {
		fmt.Println("s1 lock err: ", err)
		return
	}
	fmt.Println("s1 acquired lock")

	s2Locked := make(chan struct{})
	go func() {
		defer close(s2Locked)
		// s2等待获取锁
		if err := m2.Lock(context.Background()); err != nil {
			fmt.Println("s2 lock err: ", err)
			return
		}
	}()

	//s1 释放锁
	if err := m1.Unlock(context.Background()); err != nil {
		fmt.Println("s1 unlock err: ", err)
		return
	}
	fmt.Println("s1 released lock")

	<- s2Locked
	fmt.Println("s2 acquired lock")
}