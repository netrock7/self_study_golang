package main

import "fmt"

// 函数
//func main() {
//	r1, r2 := a(3, "条件不成立")
//	fmt.Println(r1, r2)
//}
//
//func a(data1 int, data2 string) (ret1 int, ret2 string) {
//	ret1 = data1
//	ret2 = data2
//	if data1 > 10 {
//		fmt.Println(data1)
//		//return data1, "data1大于1了，data2的内容是：" + data2
//		return
//	} else {
//		fmt.Println(data2)
//		//return data1, "data1小于1了，data2的内容是：" + data2
//		return
//	}
//	return ret1, ret2
//}

//匿名函数
//func main() {
//	b := func(data1 string) {
//		fmt.Println(data1)
//	}
//	b("我是匿名函数")
//}

//不定向函数
//func mo(data1 int, data2 ...string) {
//	fmt.Println(data1, data2)
//	for k, v := range data2 {
//		fmt.Println(k, v)
//	}
//	rr := append(data2, "555")
//	fmt.Println(rr)
//}
//func main() {
//	mo(9527, "1", "2", "3", "4")
//}

//自执行函数
//func main() {
//	(func() {
//		fmt.Println("我和别人没关系，我自己执行自己")
//	})()
//}

//闭包函数
//func main() {
//	mo()(6)
//}
//func mo() func(num int) {
//	return func(num int) {
//		fmt.Println("我被调用了，入参是：", num)
//	}
//}

//延迟调用 使用defer后不论在什么位置，都会被延迟到最后执行
func main() {
	defer mo()
	fmt.Println("等会儿")
	fmt.Println("还有一个")
	fmt.Println("没有了")
}

func mo() {
	fmt.Println("我要最开始执行")
}
