package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(regexp.MustCompile("ll").MatchString("hello"))
	fmt.Println(regexp.MustCompile("world").MatchString("hello"))
}
