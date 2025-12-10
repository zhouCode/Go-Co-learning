package main

import "fmt"

// func main() {
// 1、
// var t1 Teacher
// fmt.Println(t1)
// fmt.Println("t1:%T,%v,%q\n", t1,t1,t1)
// t1.name = "张三"
// t1.age = 30
// t1.sex = "男" 
// fmt.Println(t1)
// 2、
// t2 := Teacher{
// t2.name: "李四",
// t2.age:  35,
// t2.sex:  "男",
// }
// fmt.Println(t2)
// fmt.Println("t2:%T,%v,%q\n", t2,t2,t2)
// 3、
// t3 := Teacher{
// 	name: "王五",
// 	age:  40,
// 	sex:  "男",
// }
// t3 = Teacher{name:"赵六",age:45,sex:"男"}
// fmt.Println(t3)
// fmt.Println("t3:%T,%v,%q\n", t3,t3,t3)
// 4、
// t4 := Teacher{"张三", 30, "男"}
// fmt.Println(t4)
// fmt.Println("t4:%T,%v,%q\n", t4, t4, t4)
// emp1 := new(Emp)
// fmt.Println("emp1:%T,%v,%p\n", emp1, emp1, emp1)
// (*emp1).name = "David"
// (*emp1).age = 30
// (*emp1).sex = 1

// //语法糖写法
// emp1.name = "David"
// emp1.age = 30
// emp1.sex = 1
// fmt.Println("emp1:%T,%v,%p\n", emp1, emp1, emp1)

// }
type Emp struct {
	name, currency string
	salary         float64
}

func main() {
	emp1 := Emp{"Daniel", "$", 2000}
	emp1.printSalary()
	printSalary(emp1)
}

//方法
func (e Emp) printSalary() {
	fmt.Printf("员工姓名：%s,薪资：%s%.2f\n", e.name, e.currency, e.salary)
}

//函数
func printSalary(e Emp) {
	fmt.Printf("员工姓名：%s,薪资：%s%.2f\n", e.name, e.currency, e.salary)
}
