package main

import (
	"fmt"
)

func main() {
	s := make([]int, 1, 1)
	s = append(s, 10)
	fmt.Printf("S：%v,len：%d,cap：%d\n", s, len(s), cap(s))
	s = append(s, 20)
	fmt.Printf("S：%v,len：%d,cap：%d\n", s, len(s), cap(s))
	s = append(s, 30)
	fmt.Printf("S：%v,len：%d,cap：%d\n", s, len(s), cap(s))
	s = append(s, 40, 50, 60, 70)
	fmt.Printf("S：%v,len：%d,cap：%d\n", s, len(s), cap(s))
	fmt.Println("=========appendSlice结束==========\n")
}
