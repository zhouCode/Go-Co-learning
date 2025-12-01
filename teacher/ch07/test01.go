package main

import "fmt"

// 结构体
// 1. 结构体名不能重复
// 2. 结构体中的属性，也叫字段，必须唯一
// 3. 同类型的成员属性可以写在一行
type Teacher struct {
	name string
	age  int8
	sex  byte
}

func main() {
	//1、var声明方式实例化结构体，初始化方式为：对象.属性=值
	// var t1 Teacher
	// fmt.Println(t1)
	// fmt.Printf("t1:%T , %v , %q \n", t1, t1, t1)
	//if t1 == nil {
	//	fmt.Println()
	//}
	// t1.name = "Steven"
	// t1.age = 35
	// t1.sex = 1
	// fmt.Println(t1)
	// fmt.Println("-------------------")

	//2、变量简短声明格式实例化结构体，初始化方式为：对象.属性=值
	// t2 := Teacher{}
	// t2.name = "David"
	// t2.age = 30
	// t2.sex = 1
	// fmt.Println(t2)
	// fmt.Println("-------------------")

	//3、变量简短声明格式实例化结构体，声明时初始化。初始化方式为：属性:值 。属性:值可以同行，也可以换行。（类似map的用法）
	// t3 := Teacher{
	// 	name: "Josh",
	// 	age:  28,
	// 	sex:  1,
	// }
	// t3 = Teacher{name: "Josh2", age: 27, sex: 1}
	// fmt.Println(t3)
	// fmt.Println("-------------------")

	//4、变量简短声明格式实例化结构体，声明时初始化，不写属性名，按属性顺序只写属性值
	t4 := Teacher{"Ruby", 30, 0}
	fmt.Println(t4)
	fmt.Println("-------------------")
}
