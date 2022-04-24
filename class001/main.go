package main

import "fmt"

func main() {
	fmt.Println("hello world")
	name := "zhangmeng"
	age := 45
	b := false
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("b: %v\n", b)

	a := [...]int{1, 3, 5, 7, 9}
	fmt.Printf("a: %v\n", a)
}
