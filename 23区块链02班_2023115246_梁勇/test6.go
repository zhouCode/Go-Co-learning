package main

import "fmt"

type Student struct {
	Name string
	Age  int
}
type Emp struct {
	Name string
	Age  int
}

func main() {
	//语法糖：简便写法
	emp1 := new(Emp)
	//(*emp1).Name = "王五"
	emp1.Name = "王五"
	emp1.Age = 30
	fmt.Println(emp1)

	//var声明
	var s1 Student
	fmt.Println(s1)
	s1.Name = "张三"
	s1.Age = 18
	fmt.Println(s1)
	//变量简短声明
	s2 := Student{
		Name: "李四",
		Age:  20,
	}
	fmt.Println(s2)

}
