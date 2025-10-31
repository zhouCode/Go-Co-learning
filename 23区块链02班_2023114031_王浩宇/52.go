package main

import "fmt"
func main(){
 a:=[...]string{"USA","China","Japan"}
 b:=a
 fmt.Println(a)
 fmt.Println(b)
 b[0]="UK"
 fmt.Println(a)
 fmt.Println(b)
 b[0]="China"
 fmt.Println(a)
 fmt.Println(b)
}
