package main

import (
	"fmt"
)
func main(){
a := [...]string{"USA","Washington","DC"}
b := a
fmt.Println(a)
fmt.Println(b)
b[0] = "china"
fmt.Println(a)
fmt.Println(b)
}
