package main

// 可变参数
import "fmt"

// ...int 告诉 Go，nums 是一个可变参数
// 在函数内部，nums 的类型是 []int (一个 int 切片)
func sum(nums ...int) int {
	fmt.Print(nums, " ") // 看看 nums 是什么
	total := 0
	// 我们可以像遍历切片一样遍历它
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	// 1. 传入多个独立的参数
	fmt.Println("Sum 1:", sum(1, 2))
	fmt.Println("Sum 2:", sum(1, 2, 3, 4))

	// 2. 如果你已经有了一个切片，想把它传入可变参数
	// 你需要使用 ... 后缀来“解开”这个切片
	mySlice := []int{5, 6, 7}
	fmt.Println("Sum 3:", sum(mySlice...))
}
