package main

import "fmt"

// 数组：长度是类型的一部分
func main() {
	// 声明一个包含 3 个 int 的数组，默认值为 0
	var arr1 [3]int
	arr1[0] = 100

	fmt.Printf("值：%v, 类型：%T\n", arr1, arr1) // 输出 [100 0 0]

	// 声明并初始化一个包含 4 个 int 的数组，使用 ... 自动推断长度
	arr2 := [...]int{10, 20, 30, 40}
	fmt.Printf("值：%v, 类型: %T\n", arr2, arr2) // 输出 [10 20 30 40]

	// 下面这行代码会编译失败！
	// cannot use arr2 (type [4]int) as type [3]int in assignment
	// arr1 = arr2

	// 多维数组
	// 声明一个 2x3 的二维数组
	var matrix [2][3]int
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6

	printMatrix(matrix)
}
func printMatrix(matrix [2][3]int) {
	fmt.Println("打印矩阵:")
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}
