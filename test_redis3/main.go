package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

// var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.20.21.46:6379",
		Password: "8888",
		DB:       9,
		// Port:     6379,
	})

	getEntry, _ := rdb.HGet("/getCityByParentId", "entry").Result()
	getUrl, _ := rdb.HGet("/getCityByParentId", "url").Result()

	fmt.Println(getEntry + getUrl)

}
