package main

import "fmt"

func main() {
	res,err3 := checkAge(-12)
	if err3 != nil {
		fmt.Println(err3.Error())
		fmt.Println(err3)
	}else{
		fmt.Println(res)
	}
}
func checkAge(age int) (string,error) {
	if age < 0 {
		err := fmt.Errorf("您输入的年龄是：%d，该数值为负数，有错误！",age)
		return "",err
	}else {
		return fmt.Sprintf("您输入的年龄是：%d",age),nil
	}
}