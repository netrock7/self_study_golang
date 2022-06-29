package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("context")

	//1. WithCancel
	// ctx, cancel := context.WithCancel(context.Background())

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	// go TestCancel(ctx)
	go TestDeadline(ctx)
	<-time.After(time.Second * 10)
	cancel()
	<-time.After(time.Second * 2)
	fmt.Println("main结束了...")
}

func TestCancel(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("主协程取消了，我是子协程也同时结束了...")
			return
		default:
			fmt.Println("我是子协程，我在运行中...")
			time.Sleep(time.Second)
		}
	}
}

func TestDeadline(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("我是子协程，我运行了5秒就结束了...")
			return
		default:
			fmt.Println("我是deadline子协程，我在运行中...")
			time.Sleep(time.Second)
		}
	}
}
