package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.ParseInt("123", 10, 64))
	fmt.Println(strconv.ParseInt("123x", 10, 64))
	fmt.Println(strconv.ParseInt("0", 10, 64))
	fmt.Println(strconv.FormatInt(123, 10))
	fmt.Println(strconv.FormatInt(0, 10))
}
