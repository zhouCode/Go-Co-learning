package main

import "fmt"

func main() {
	// //写法1
	// score := 85
	// if score > 60 {
	// 	fmt.Println("及格")
	// } else {
	// 	fmt.Println("不及格")
	// }
	// //写法2:在if语句中初始化
	// if score2 := 85; score2 > 60 {
	// 	fmt.Println("及格")
	// } else {
	// 	fmt.Println("不及格")
	// }

	for i := 0; i < 10; i++ {	//三个条件都可省略
		fmt.Println(i)
	}
}
