package main

import "fmt"

// 数组
//func main() {
//	a := [3]int{1, 2, 3}
//	b := [...]int{123, 4234, 34534, 6546, 7547, 578, 7689, 769, 7890, 890}
//	var d = new([10]int)
//	d[5] = 8
//	fmt.Println(a, b, d)
//
//	zoo := [...]string{"鸭子", "鲨鱼", "螃蟹"}
//	for i := 0; i < len(zoo); i++ {
//		fmt.Println(zoo[i] + "站起来了")
//	}
//
//	for i, v := range zoo {
//		fmt.Println(i, v)
//	}
//
//	fmt.Println(len(zoo), cap(zoo))
//}

//func main() {
//	a := [3]int{0, 1, 2}
//	c1 := a[1:]
//	fmt.Println(a, c1)
//	//a = append(a, 3)
//	//fmt.Println(a, c1)
//	c1 = append(c1, 5, 8)
//	c1[0] = 4
//	fmt.Println(a, c1)
//}

// cap会自动扩充容量，len根据切片实际内容增减
//func main() {
//	a := [3]int{0, 1, 2}
//	cl := a[:]
//	cl = append(cl, 5)
//	cl = append(cl, 5)
//	cl = append(cl, 5)
//	cl = append(cl, 5)
//
//	fmt.Println(len(cl), cap(cl))
//}

func main() {
	a := [3]int{0, 1, 2}
	cl := a[:]
	cl1 := a[2:]
	fmt.Println(cl1)
	cl = append(cl, 5)
	fmt.Println(cl)
	copy(cl, cl1)
	fmt.Println(cl)

	var aa []int

	aaa := make([]int, 4)

	fmt.Println(aa, aaa)
}
