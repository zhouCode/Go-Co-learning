package main

import "fmt"

//定义 Teacher 结构体
type Teacher struct {
	name string
	age  int
	sxe  int
}

func main() {
	//1.var声明方式实例化结构体，初始化方式为：对象.属性=值
	var t1 Teacher
	fmt.Println(t1)
	fmt.Printf("t1:%T ,%v , %q\n", t1, t1, t1)
	t1.name = "Steven"
	t1.age = 35
	t1.sxe = 1
	fmt.Println(t1)
	fmt.Println("-----------------")

	//3.变量简短声明格式实例化结构体，声明时初始化，初始化方式为：属性:值,属性：值可以同行
	t3 := Teacher{
		name: "John",
		age:  30,
		sxe:  0,
	}
	t3 = Teacher{name: "John", age: 30, sxe: 0}
	fmt.Println(t3)
	fmt.Printf("-----------------")

	//4.变量简短声明格式实例化结构体，声明时初始化,不写属性名，按属性顺序只写属性值
	t4 := Teacher{"Ruby", 28, 0}
	fmt.Println(t4)
	fmt.Printf("-----------------")
}
