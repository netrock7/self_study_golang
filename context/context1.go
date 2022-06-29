package main

import (
	"fmt"
	"time"
)

func main1() {
	fmt.Println("context")

	ch := make(chan bool)
	// 1. runtime.Goexit()
	// 2. channel
	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("我要结束了...")
				return
			default:
				fmt.Println("我是协程...")
				time.Sleep(time.Second)
				// runtime.Goexit()
			}
		}
	}()

	<-time.After(time.Second * 5)
	ch <- true

	fmt.Println("main结束")
}
