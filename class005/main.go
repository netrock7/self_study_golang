package main

import "fmt"

// 数组
func main() {
	a := [3]int{1, 2, 3}
	b := [...]int{123, 4234, 34534, 6546, 7547, 578, 7689, 769, 7890, 890}
	var d = new([10]int)
	d[5] = 8
	fmt.Println(a, b, d)

	zoo := [...]string{"鸭子", "鲨鱼", "螃蟹"}
	for i := 0; i < len(zoo); i++ {
		fmt.Println(zoo[i] + "站起来了")
	}

	for i, v := range zoo {
		fmt.Println(i, v)
	}

	fmt.Println(len(zoo), cap(zoo))
}
