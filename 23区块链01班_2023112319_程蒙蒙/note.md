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


