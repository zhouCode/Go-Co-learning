package main

import "fmt"

func main(){
	a := 10
	var ip *int = &a
	fmt.Printf("ip=%p \n",ip)
	fmt.Printf("*ip=%d \n",*ip)
	fmt.Printf("a=%d \n",&a)
}