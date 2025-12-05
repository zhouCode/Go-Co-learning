# 程蒙蒙 的学习笔记

## 基本信息
- **学号**: 2023112319
- **班级**: 23区块链01班
- **姓名**: 程蒙蒙

## 学习记录

### 课程笔记
- 区块链基础
- 智能合约
-2025.10.10
- 合约基础

- iotet特殊常量值,只能在常量赋值中出现
- const (
    a = intet
    b
    c
)
- 常量赋值时,可以使用iota特殊常量值,从0开始,每次赋值自动加1
- 在每一个const关键字出现时，重新定义为0
- 2025.10.17
- 第三章Go语言流程控制
if 布尔表达式 {
- 条件没有小括号
- 必须要有大括号
}
首字母大写，代表在外部可以调用
左括号必须和if或else同行
switch varl {
    case val1 ://执行每条case自动break
    ...//fallthrough实现穿透情况，强制执行下一个case模块
    case val2:
    ...
    default:
    ...
}    
for循环
for 初始化语句; 布尔表达式; 迭代语句 {
    循环体
}
for 一个条件{
    语句1
}//相当于while循环

for{
    语句一
}//无限循环

OuterLoop:
for{
    for{
        语句二
        break OuterLoop
    }
}//有OuterLoop标签的for循环,可以在内部循环中使用break语句跳出外部循环（使用break+OuterLoop,直接跳出到OuterLoop那一层）

- 第四章
函数
- 函数定义
func 函数名(参数列表) (返回值列表) {
    函数体
}
func add (a,b int){}

func myfunc(arg... int) {}
- error:错误类型
- nil:代表空，表示error不存在
fun abhf(a,b int) (return int, remainder int){
    a+b
    a%b
    return
}
- 函数变量
- var f func((int,int) int)//声明一个变量f，他的类型是一个函数，这个函数有两个int类型的参数，返回一个int类型的值
- 函数调用
函数名(参数列表)
- 函数参数
- 函数返回值
- 函数类型
- 函数作为参数
- 函数作为返回值

- 匿名函数
- 匿名函数定义
func(参数列表) (返回值列表) {
    函数体
}
- 匿名函数调用
func(参数列表)(返回值列表){
    函数体
}(参数列表)
func (data int) {
    fmt.Println("hello",data)
}
- 闭包=函数+对其外部作用域的引用（引用环境）
- 闭包定义
func 函数名(参数列表) (返回值列表) {
    函数体
}
- 闭包调用
函数名(参数列表)
- 闭包示例
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}
- 闭包作用
- 闭包可以访问和操作其外部函数的变量
- 闭包可以用于实现函数柯理化
- 闭包可以用于实现函数工厂
- 闭包可以用于实现函数记忆化

- 可变参数
fmt.Println("Sum 3:",sum(myAlice...))//需要用...来解包切片
-递归函数
func factorial(n int) int {
    //退出条件！！！
    if n == 0 {
        return 1
    }
    //递归调用
    return n * factorial(n-1)
}
- 指针
- 指针定义
var p *int//p是一个指向int类型的指针
- 指针赋值
var a int = 10
var p *int = &a//p指向a的地址
- 指针解引用
var a int = 10
var p *int = &a//p指向a的地址
fmt.Println(*p)//输出a的值
- 指针作为参数
func changeValue(p *int) {
    *p = 100//通过指针修改指向变量的值
}
- 指针作为返回值
func getPointer() *int {
    var a int = 10
    return &a//返回a的地址
}
- 1,对于一个变量，想获取他的地址，就使用&[变量名]
- 2,对于一个指针变量，想获取他指向的变量，就使用*[指针变量名]
- 3,声明一个指针变量，使用*[类型]
  - |--------- 第五章 ---------|
- 数组定义
var 变量名 [数组长度] 数组类型
var a [5]int//a是一个有5个int类型元素的数组
不能把长度不一致的数组作为参数传递（例：a[3]和a[5]不能相互赋值）,数组是值类型
- 数组初始化
var a [5]int = [5]int{1,2,3,4,5}//a是一个有5个int类型元素的数组，元素值分别是1,2,3,4,5
- 数组遍历
方法一：
for i := 0; i < len(a); i++ {
    fmt.Println(a[i])
}
方法二：
for _,v := range a{
    fmt.Println(v,"\t")
}
- 数组作为参数
func sum(arr [5]int) int {
    sum := 0
    for i := 0; i < len(arr); i++ {
        sum += arr[i]
    }
    return sum
}
- 切片：
 变长数组，每个元素类型都相同
var 切片名 []type
var slicel []type = make([]type, len)
var s []int//s是一个int类型的切片
- 切片初始化
var s []int = []int{1,2,3,4,5}//s是一个有5个int类型元素的切片，元素值分别是1,2,3,4,5
s := [] int {1,2,3}
//左闭右开区间，包含1,2,3不包含4
s := [] int {1,2,3,4,5}
s := s[1:4]//s是一个有3个int类型元素的切片，元素值分别是2,3,4
//左闭右开区间，包含1,4不包含5
s := s[1:4]//s是一个有3个int类型元素的切片，元素值分别是2,3,4
- 切片遍历
方法一：
for i := 0; i < len(s); i++ {
    fmt.Println(s[i])
}
方法二：
for _,v := range s{
    fmt.Println(v,"\t")
}
- 切片作为参数
func sum(s []int) int {
    sum := 0
    for i := 0; i < len(s); i++ {
        sum += s[i]
    }
    return sum
}
append(向谁追加，追加的元素（可以多个）：向切片追加元素，需要用变量接收返回值）
s := []int{1,2,3}
s = append(s,4,5,6)
fmt.Println(s)//[1 2 3 4 5 6]
len长度到达一定数量时，会自动扩容，扩容后的容量是原来的2倍

- map是无序的，每次遍历的顺序都可能不同
- map的key必须是可比较的类型，不能是函数、切片、映射
- map的value可以是任意类型
- map的key-value对是无序的
- map中key必须是唯一的

- |---------第六章 --------|
string: 字符串
range遍历:
....
len(str):字符串的长度，中文占3个字节，英文占1个字节
Contains:判断字符串是否包含指定的子字符串
func Contains(str, substr string) bool
substr:子字符串
    返回值：如果str包含substr，则返回true，否则返回false
分割字符串：
Split:将字符串按照指定的分隔符进行分割，返回一个字符串切片
func Split(str, sep string) []string
sep:分隔符
    返回值：一个字符串切片，包含分割后的子字符串
ToLower:将字符串中的所有字符转换为小写
ToUpper:将字符串中所有字符转换为大写
TrimSpace:修剪字符串首尾的空格
- strconv包中常用函数：
ParseInt:将字符串转换为整数
func ParseInt(s string, base int, bitSize int) (i int64, err error)
s:要转换的字符串
base:进制数，2<=base<=36
bitSize:结果的位大小，0表示根据字符串的内容自动选择位大小
    返回值：转换后的整数和错误信息
regexp:正则表达式
- 邮箱验证正则表达式示例
```go
package main

import (
    "fmt"
    "regexp"
)

func main() {
    // 邮箱正则表达式模式
    emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
    
    // 测试邮箱地址
    testEmails := []string{
        "chengmm2023@example.com",      // ✅ 有效
        "student_123@university.edu", // ✅ 有效
        "test.email+tag@domain.co",   // ✅ 有效
        "invalid.email@",             // ❌ 无效
        "@invalid.com",                // ❌ 无效
        "noatsign.com",                // ❌ 无效
        "spaces in@email.com",         // ❌ 无效
    }
    
    // 编译正则表达式
    re := regexp.MustCompile(emailRegex)
    
    fmt.Println("📧 邮箱验证测试:")
    for _, email := range testEmails {
        if re.MatchString(email) {
            fmt.Printf("✅ 有效邮箱: %s\n", email)
        } else {
            fmt.Printf("❌ 无效邮箱: %s\n", email)
        }
    }
    
    // 提取文本中的邮箱
    text := "我的邮箱是 chengmm@example.com，备用邮箱是 backup123@domain.org.cn"
    emailPattern := `[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`
    re2 := regexp.MustCompile(emailPattern)
    emails := re2.FindAllString(text, -1)
    fmt.Printf("\n从文本提取的邮箱: %v\n", emails)
}
```

- 正则表达式语法说明：
  - `^` 匹配字符串开始
  - `[a-zA-Z0-9._%+-]+` 匹配用户名部分（允许字母、数字、点、下划线、百分号、加号、减号）
  - `@` 匹配@符号
  - `[a-zA-Z0-9.-]+` 匹配域名部分
  - `\.` 匹配点号（需要转义）
  - `[a-zA-Z]{2,}` 匹配顶级域名（至少2个字母）
  - `$` 匹配字符串结束

 Round:四舍五入到最近的整数
 max：最大
 min：最小
 math:绝对值
 Pow:平方根   
randNum := rand.Intn(100)//生成0到99的随机数    


|---------------第七章----------------|
go语言面向对象编程
特点：封装，继承，多态
定义一个结构体：
type 类型名 struct {
    成员属性1 类型1
    成员属性2 类型2
    成员属性3 类型3
    成员属性4 类型4
}
同一个包里面，结构体名字不能重复
同类型的属性可以写在一行
var声明方式实例化结构体，初始化方式为：对象.属性=值
var  变量名 结构体名
变量名.属性1 = 值1
变量名.属性2 = 值2
变量名.属性3 = 值3
.......
用new()实例化类型是一个指针

函数
func 函数名（参数列表） （返回值类型） {
    //函数体
    return 返回值
}



|-------------第8章------------------|
error接口：
type error interface {
    Error() string
}error本质上是一个接口类型

error常用返回值来表示错误，通常是返回值的最后一个参数
自定义error：实现error方法
type MyError struct {
    Msg string
}
func (e *MyError) Error() string {
    return e.Msg
}

errors.New:创建一个新的错误
func New(text string) error
text:错误信息
    返回值：一个error接口类型的错误

fmt.Errorf:格式化错误信息(直接返回一个错误类型并输出具体参数)
func Errorf(format string, a ...any) error
format:格式化字符串
a:格式化参数
    返回值：一个error接口类型的错误

defer:延迟一个函数或方法的执行
特性：
    1. 延迟的函数或方法会在包含它们的函数或方法返回之前执行
    2. 多个defer语句会按照后进先出（LIFO）的顺序执行
    3. 延迟的函数或方法可以访问包含它们的函数或方法的参数和局部变量
### |-----------------第12章 -----------------|
- Go语言并发编程
并发：多个任务在同一时间段内执行
并行：多个任务在同一时刻执行
- 线程和进程
在操作系统下，线程是独立运行和独立调度的基本单位
一个进程可以包含多个线程，线程是容器中的工作单位
协程执行的任务单位比线程小，一个线程可以包含多个协程
协程是编译器级别的，由编译器负责调度和切换，而线程是操作系统级别的，由操作系统负责调度和切换
- goroutine属于抢占式任务处理。
匿名函数启动goroutine：
func main() {
    go func() {
        fmt.Println("hello goroutine")
    }()
    fmt.Println("main goroutine")
}
channel:引用类型
语法
var channel 变量 chan channel 类型
实现接口下所有方法就代表实现了接口
旧版：interface{}
新版：any
ch2 := make(chan interface{})
等价于
ch2 := make(chan any)
使用channel发送数据
语法：channel 变量 <- 值

data , ok := <-channel变量
    如果channel变量中没有数据，ok为false
    如果channel变量中没有数据，ok为true
OK：channel是否被关闭

可以从关闭的channel中读取数据
time：
Time包
time.Now()：返回当前时间
time.Sleep()暂停执行指定时间
    参数：时间 duration
    返回值：无


After()函数：
    功能：返回一个时间通道，在指定时间后发送当前时间
    参数：时间 duration
    返回值：一个时间通道

select语句：
    select分支语句：
        ase 通道变量 <- 值：
            当通道变量可以发送数据时，执行该分支
        case 值 := <- 通道变量：
            当通道变量可以接收数据时，执行该分支
        default：//不需要等待
            当所有通道变量都不能发送或接收数据时，执行该分支
sync包
WaitGroup：等待一组 goroutine 完成
    方法：
        Add(delta int)：添加 delta 个 goroutine 到 WaitGroup 中
        Done()：当前 goroutine 完成，将 WaitGroup 中的计数器减一
        Wait()：阻塞当前 goroutine，直到 WaitGroup 中的计数器为零
        
    func (wg *WaitGroup) Add(delta int)
        delta:要添加的 goroutine 数量
        注意：delta 可以是负数，用于减少 goroutine 数量
互斥锁：Mutex
    方法：
        Lock()：加锁
        Unlock()：解锁
        e.g.多个售票窗口同时售票，每个窗口售票时间不同，需要互斥锁来防止多个窗口同时售票
    注意：
        1. 互斥锁只能在一个 goroutine 中加锁，其他 goroutine 必须等待解锁后才能加锁
        2. 互斥锁加锁后，其他 goroutine 必须等待解锁后才能加锁，否则会导致死锁
读写锁：RWMutex
    方法：
        Lock()：加写锁
        Unlock()：解锁写锁
        RLock()：加读锁
        RUnlock()：解锁读锁
    注意：
        1. 读写锁可以同时多个 goroutine 加读锁，但是只能有一个 goroutine 加写锁
        2. 读写锁加写锁后，其他 goroutine 必须等待解锁后才能加锁，否则会导致死锁
        3.写操作是互斥的，读和写是互斥的，读和读不互斥
条件变量：Cond
    方法：
        Wait()：阻塞当前 goroutine，直到条件变量满足
        Signal()：唤醒一个等待的 goroutine
        Broadcast()：唤醒所有等待的 goroutine
    注意：
        1. 条件变量必须与互斥锁配合使用
        2. 条件变量的 Wait() 方法会自动解锁互斥锁，在返回前会重新加锁互斥锁

    
### 作业记录
- 完成区块链基础概念学习
- 实现简单的智能合约示例
- 学习Solidity语法基础
-   学习
### 项目经验
- 

### 学习心得
- 

---
*创建时间: 2025-09-25 08:57:45*








