package main

import "fmt"
func main(){
	var arr1 [5]int
	arr1[0] = 100
	arr1[1] = 200
	arr1[2] = 300
	arr1[3] = 400
	arr1[4] = 500
	fmt.Printf("arr1=%v \n",arr1)
	arr2 := [5]int{1,2,3,4,5}
	fmt.Printf("arr2=%v \n",arr2)
}	