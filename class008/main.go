package main

import "fmt"

//指针和地址
//指针
//func main() {
//	var a int
//	a = 123
//	fmt.Println(a)
//	var b *int	//声明b为一个指针
//	b = &a	//获取a的地址赋值给b
//	*b = 456	//给b指针指向的地址赋值，这个时候b的指针指向的就是a变量所在的位置
//	fmt.Println(a, b, *b)
//	fmt.Println(a == *b, &a == b)	//可以印证，b指针指向的就是a变量，a变量所在的内存地址就是b指针
//}

// 数据指针和指针数组
func main() {
	var arr [5]string
	arr = [5]string{"1", "2", "3", "4", "5"}
	var arrP *[5]string
	arrP = &arr
	fmt.Println(arr, arrP)

	var arrpp [5]*string

	var str1 string = "str1"
	var str2 string = "str2"
	var str3 string = "str3"
	var str4 string = "str4"
	var str5 string = "str5"

	arrpp = [5]*string{&str1, &str2, &str3, &str4, &str5}
	*arrpp[2] = "666"
	fmt.Println(str3)
	fmt.Println(*arrpp[0], *arrpp[1], *arrpp[2], *arrpp[3], *arrpp[4])
}
