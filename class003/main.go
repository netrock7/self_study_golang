package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 整形变量
	var num1 uint
	var num2 int
	num1 = 999
	num2 = -999
	fmt.Println(num1, num2)
	fmt.Printf("%T", num2)

	// 浮点数类型
	var num3 float64
	num3 = 2332.34224243
	fmt.Println(num3)
	fmt.Printf("%T", num3)

	// 字符串类型
	var str1 string
	str1 = "34"
	fmt.Println(str1)
	fmt.Printf("%T", str1)
	int, _ := strconv.Atoi(str1)
	fmt.Println(int)
	fmt.Printf("%T", int)

	// 布尔类型
	var bool1 bool
	bool1 = false
	fmt.Println(bool1)
	fmt.Printf("%T", bool1)
}
