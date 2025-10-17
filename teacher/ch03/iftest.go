package main

import "fmt"

func main() {
	score := 85
	// 不同点 1：条件没有小括号 ()
	// 不同点 2：必须有大括号 {}
	// if (score>60)
	//     System.out.println("及格了");
	//
	if score > 60 {
		fmt.Println("及格了")
	} else {
		fmt.Println("不及格")
	}
	// 不同点 3 (最重要)：支持 if 内初始化语句
	if score2 := 85; score2 > 60 {
		fmt.Println("及格了")
	} else {
		fmt.Println("不及格")
	}

}
