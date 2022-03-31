package main

import "fmt"

func main() {
	// go没有后增，只有前增
	//a := 6
	//a++
	//fmt.Println(a)

	// 条件判断 if else
	//if a > 3 {
	//	fmt.Println("正确")
	//} else if a == 1 {
	//	fmt.Println("是一")
	//} else {
	//	fmt.Println("错误")
	//}

	// 条件判断 switch
	//switch a {
	//case 0:
	//	fmt.Println(a)
	//case 1:
	//	fmt.Println(a)
	//case 2:
	//	fmt.Println(a)
	//default:
	//	fmt.Println("nothing")
	//}

	// 条件判断 for循环
	//for {
	//	a++
	//	fmt.Println(a)
	//	if a > 100 {
	//		break
	//	}
	//}

	//for b := 0; b < 100; b++ {
	//	fmt.Println(b)
	//}

	// 循环体，跳转，goto指定循环体
	a := 0
A:
	for a < 10 {
		a++
		fmt.Println(a)
		if a == 8 {
			break A
			goto B
		}
	}
B:
	fmt.Println("welcome to B area")
}
