package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	fmt.Println("随机数:", r.Intn(100))
	fmt.Println("随机浮点数:", r.Float64())
}
