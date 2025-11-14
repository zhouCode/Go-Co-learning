package main

import (
	"fmt"
)

func main() {
	emp1 := new(Emp)
	fmt.Printf("emp1:%T,%v,%q\n", emp1, emp1, emp1)

	(*emp1).Name = "张三"
	(*emp1).Age = 18
	(*emp1).Sex = 1
	fmt.Printf("emp1:%T,%v,%q\n", emp1, emp1, emp1)

}
