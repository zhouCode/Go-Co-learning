package main

import (
	"fmt"
)

func main() {
	// var t1 Teacher
	// fmt.Println(t1)
	// fmt.Printf("t1:%T,%v,%q\n", t1, t1, t1)
	// t1.Name = "张三"
	// t1.Age = 18
	// fmt.Printf("t1:%T,%v,%q\n", t1, t1, t1)

	// t3 := Teacher{
	// 	Name: "李四",
	// 	Age:  20,
	// 	sex:1,
	// }
	// t3 = Teacher{name:"王五",Age:21,sex:1}
	// fmt.Printf("t3:%T,%v,%q\n", t3, t3, t3)

	t4 := Teacher{
		Name: "赵六",
		Age:  22,
		sex:  1,
	}
	fmt.Printf("t4:%T,%v,%q\n", t4, t4, t4)
}
