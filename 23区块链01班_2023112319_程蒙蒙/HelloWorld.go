//package main

//import "fmt"

// func main() {
// 	var a int = 10
// 	var b = 10
// 	c := 20 //第一次赋值才能用
// 	fmt.Println("hello world:", a, b, c)
// }

//package main
//import "fmt"
// func main() {
// 	a := 3
// 	switch a {
// 	case 1:
// 		fmt.Println("a=1")

// 	case 2:
// 		fmt.Println("a=2")
// 	default:
// 		fmt.Println("default")
// 	}
// }

//package main
//import "fmt"
// func main() {
//     fmt.Println("Hello, World!")
//     fmt.Println("欢迎 程蒙蒙 学习Go语言!")
//     fmt.Println("这是你的第一个Go程序")

//     // 基本变量声明
//     var name string = "程蒙蒙"
//     var age int = 20

//     fmt.Printf("学生姓名: %s\n", name)
//     fmt.Printf("年龄: %d\n", age)

//     // 简单的函数调用
//     greet(name)
// }

// // 问候函数
// func greet(name string) {
//     fmt.Printf("你好, %s! 欢迎来到Go语言的世界!\n", name)
// }

// package main

// import "fmt"

// func main() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(i)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	for {
// 		fmt.Println("hello world")
// 	}
// }

// package main

// import "fmt"

// func main() {
// OuterLoop:
// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 3; j++ {
// 			fmt.Printf("正在检查: i=%d, j=%d\n", i, j)
// 			if i == 1 && j == 1 {
// 				fmt.Println("找到了 跳出所有循环！")
// 				break OuterLoop
// 			}
// 		}
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	multiply := func(a,b int) int {//multiply是一个函数变量，他的类型是func(int,int) int
// 		return a*b
// 	}
// 	fmt.Println("匿名函数相乘：",multiply(2,3))

// 	//定义并立即执行
// 	//最后的（5,3）是在调用它
// 	result := func(a,b int) int {//result是int类型
// 		return a*b
// 	}(5,3)
// 	fmt.Println("立即执行结果：",result)
// }


// package main
// import "fmt"

// func intSequence() func() int {
// 	i :=0
// 	return func() int {
// 		i++
// 		return i
// 	}
// }
// func main() {
// 	nextInt := intSequence()
// 	fmt.Println(nextInt())
// 	fmt.Println(nextInt())
// 	fmt.Println(nextInt())

// 	newInts := intSequence()
// 	fmt.Println(newInts())
// 	fmt.Println(newInts())
// 	fmt.Println(nextInt())

// }


// package main
// import "fmt"

// func main(){
// 	a := 10
// 	var ip *int = &a
// 	fmt.Println("a 现有指向变量是:%x\n",&a)
// 	fmt.Println("ip 现有指向变量的值是:%v\n",*ip)
// 	//1,对于一个变量，想获取他的地址，就使用&[变量名]
// 	//2,对于一个指针变量，想获取他指向的变量，就使用*[指针变量名]
// 	//3,声明一个指针变量，使用*[类型]
// }
//第五章:数组
// package main
//  import "fmt"

//  func main() {
// 	var arr1 [3]int
// 	arr1[0] =100
// 	fmt.Printf("值：%v,类型：%T\n",arr1,arr1)
// 	arr2 := [...]int{10,20,30,40}
// 	fmt.Printf("值：%v,类型：%T\n",arr2,arr2)
//  }

//  package main
//  import "fmt"

//  func main() {
// 	a :=[...]string{"USE","CHINA","JAPAN"}
// 	b :=a
// 	fmt.Println("修改前:")
// 	fmt.Println("a:",a)
// 	fmt.Println("b",b)
// 	b[0] = "Singapore"
// 	fmt.Println("修改后:")
// 	fmt.Println("a:",a)
// 	fmt.Println("b",b)
//  }

//   package main
//  import "fmt"

//  func main() {
// 	s := make([] string,3,5)
// 	parent := [...]int{10,20,30,40,50}
// 	child := parent[1:4]
// 	fmt.Println("Slice:",s)
// 	fmt.Println("len:",len(s))
// 	fmt.Println("Cap:", cap(s))
// 	fmt.Println("child:",child)
// 	child[0] = 99
// 	fmt.Println("After modify child,Parent",parent)
// }

//   package main
//  import "fmt"

//  func main() {
// 	s :=make([]int,1,1)
// 	s = append(s,10)
// 	fmt.Printf("s: %v, len: %d, cap: %d\n", s, len(s), cap(s))
// 	s = append(s,20)
// 	fmt.Printf("s: %v, len: %d, cap: %d\n",s,len(s),cap(s))
// 	s = append(s,30)
// 	fmt.Printf("s: %v, len: %d, cap: %d\n",s,len(s),cap(s))
// 	s = append(s,40)
// 	fmt.Printf("s: %v, len: %d, cap: %d\n",s,len(s),cap(s))
// 	fmt.Printf("=======appendSlice 结束 =======\n")
// 	}

//第六章
//  package main
//  import "fmt"

//  func main() {
// 	s :="我爱go语言"
// 	for _,value  := range s {
// 		fmt.Printf("%c",value)
// 	}
// 	fmt.Println()
// 	fmt.Printf("%d",len(s))
// 	for i,ch := range []byte(s){
// 		fmt.Printf("%d:%x",i,ch)
// 	}
// 	fmt.Println()
// }


// package main

// import (
// 	"fmt"
// 	"strconv"
// )

//  func main() {
// 	fmt.Println(strconv.ParseInt("123",10,64))
// 	fmt.Println(strconv.ParseInt("0",10,64))
// 	fmt.Println(strconv.ParseInt("123x",10,64))

// 	fmt.Printf("%T \n",strconv.FormatInt(123,10))
// 	fmt.Println("%T",strconv.FormatInt(0,10))
// }


// package main

// import (
// 	"fmt"
// 	"time"
// )

//  func main() {
// 	t, _ := time.Parse("2006-01-02 15:04:05","2023-08-01 12:34:56")
// 	fmt.Println("解析时间： ",t)
// 	now :=time.Now()
// 	fmt.Println("当前时间:", now.Format("2006/01/02 15:04:05"))
// 	fmt.Println("24小时后的时间: ",now.Add(24*time.Hour))
// 	fmt.Println("时间差: ", now.Sub(t))
// 	fmt.Println("解析时间是否在当前时间之前: ",t.Before(now))
// 	fmt.Println("解析时间是否在当前时间之后: ",t.After(now))
// }

// package main
// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )
// func main(){
// 	seed := rand.NewSource(time.Now().UnixNano())
// 	r1 := rand.New(seed)
// 	fmt.Println("随机数:", rand.Intn(100))
// 	fmt.Println("随机数:",r1.Intn(100))
// 	fmt.Println("随机浮点数: ",r1.Float64())
// }



