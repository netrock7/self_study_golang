package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.8.123:6379",
		Password: "8888",
		DB:       0,
		// Port:     6379,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("连接redis出错，错误信息： %v", err)
	}
	fmt.Println("成功连接redis")

	keys, err := rdb.Keys(ctx, "*").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(keys)

	vType, err := rdb.Type(ctx, "111").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(vType) //string
}
