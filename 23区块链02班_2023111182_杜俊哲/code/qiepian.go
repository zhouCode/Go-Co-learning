package main

import (
	"fmt"
)
func main(){
s:= make([]string,3,5)
parent := [...]int{10,20,30,40,50}
child := parent[1:4]
fmt.Println("slice:",s)
fmt.Println("cap:",cap(s))
fmt.Println("len:",len(s))
fmt.Println(child)
}
