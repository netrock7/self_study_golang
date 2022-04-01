package main

import (
	"fmt"
)

// 数据类型map
//func main() {
//	var countryCapitalMap map[string]string /*创建集合 */
//	countryCapitalMap = make(map[string]string)
//
//	/* map插入key - value对,各个国家对应的首都 */
//	countryCapitalMap["France"] = "巴黎"
//	countryCapitalMap["Italy"] = "罗马"
//	countryCapitalMap["Japan"] = "东京"
//	countryCapitalMap["India "] = "新德里"
//
//	/*使用键输出地图值 */
//	for country := range countryCapitalMap {
//		fmt.Println(country, "首都是", countryCapitalMap[country])
//	}
//}

//func main() {
//	var m map[string]string
//	m = map[string]string{} //必须要初始化
//	fmt.Println(m)
//	m["name"] = "solo"
//	m["sex"] = "man"
//	fmt.Println(m)
//}

func main() {
	//m1 := map[int]bool{}
	//m1[1] = false
	//m1[2] = true
	//fmt.Println(m1)

	m1 := map[int]interface{}{}
	m1[1] = 10
	m1[2] = true
	m1[3] = "success"
	m1[4] = []int{1, 2, 3}
	fmt.Println(m1, len(m1))

	//delete(m1, 2)
	//fmt.Println(m1, len(m1))

	for k, v := range m1 {
		fmt.Println(k, v)
	}
}
