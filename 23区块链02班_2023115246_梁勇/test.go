package main

import "fmt"

func main() {
	score := 90
	if score > 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}
	//if第二种方法
	if score2 := 85; score2 > 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

	//switch语句
	switch score3 := 85; score3 {
	case 90:
		fmt.Println("A")
	case 80:
		fmt.Println("B")
	case 70:
		fmt.Println("C")
	case 60:
		fmt.Println("D")
	default:
		fmt.Println("default")
	}

	switch score4 := 85; score4 {
	case 85:
		fmt.Println("1")
		//加入fallthrough后，会继续执行下一个case,这里会输出1和2
		fallthrough
	case 80:
		fmt.Println("2")
		fallthrough
	default:
		//加入fallthrough后，会继续执行下一个case,这里会输出default
		fmt.Println("default")
	}

	//模式1；经典for循环（类似C/C++/Java）
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//模式2；替换while
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	//模式3；无限循环（类似Python的while True）
	// for {
	// 	fmt.Println("无限循环")
	//}
	// 	break
	//模式4；for-range循环（类似Python）a=97
	for index, value := range "hello" {
		fmt.Println(index, value)
	}
}
2025.10.17;14:49周五
