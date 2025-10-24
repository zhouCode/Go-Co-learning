package main

import "fmt"

func main(){
	pos := adder()
	for i := 0; i < 5; i++ {
		fmt.Printf("i=%d \t",i)
		fmt.Printf("pos=%d \n",pos(i))
	}
	fmt.Println("--------------------------------")
	for i := 0; i < 5; i++ {
		fmt.Printf("i=%d \t",i)
		fmt.Printf("pos=%d \t",pos(i))
	}
}
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		fmt.Printf("sum=%d \t",sum)
		sum += x
		fmt.Printf("sum=%d \t",sum)
		return sum
	}
}