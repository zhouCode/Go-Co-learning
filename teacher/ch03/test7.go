package main

import "fmt"

func main() {
	// 遍历切片 (slice)
	nums := []int{10, 20, 30, 40, 50}
	// range 会返回两个值：索引(index) 和 值(value)
	for index, value := range nums {
		fmt.Printf("索引: %d, 值: %d\n", index, value)
	}

	// 如果你只关心值，不关心索引，可以用下划线 _ 忽略它
	fmt.Println("--- 只关心值 ---")
	for _, value := range nums {
		fmt.Println("值:", value)
	}

	// 遍历 map
	// capitals := map[string]string{"China": "北京", "Japan": "东京", "USA": "华盛顿"}
	// for country, capital := range capitals {
	// 	fmt.Printf("%s 的首都是 %s\n", country, capital)
	// }
}
