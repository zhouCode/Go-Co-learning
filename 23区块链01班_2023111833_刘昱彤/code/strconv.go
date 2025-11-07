package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.ParseInt("123", 10, 64))
	fmt.Println(strconv.ParseFloat("3.14", 64))
	fmt.Println(strconv.FormatInt(123, 10))
}
