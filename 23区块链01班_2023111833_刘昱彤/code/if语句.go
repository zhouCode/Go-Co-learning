package main

func main() {
	score := 85 //: 初始化，第一次赋值的时候使用
	if score > 60 {
		println("及格")
	} else {
		println("不及格")
	}
	if score2 := 90; score2 > 60 {
		println("及格")
	} else {
		println("不及格")
	}
}
