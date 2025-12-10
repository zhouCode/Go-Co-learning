package main

import "fmt"

// --- 任务一：结构体与方法 ---
/*
要求：
1. 定义 Dog 结构体，包含 Name (string) 和 Energy (int)。
2. 定义 Bark() 方法，打印 "汪汪"。
3. 定义 Eat() 方法，使 Energy 增加 10。
*/
// TODO 1: 定义 Dog 结构体
type Dog struct {
	Name string
	Energy int
}

// TODO 2: 为 Dog 定义 Bark 方法
// 提示: func (d Dog) Bark() { ... }
func (d Dog) Bark() {
	// 打印 "Dog [名字] is barking: 汪汪!"
	fmt.Printf("Dog %s is barking：汪汪！\n",d.Name)
}

// TODO 3: 为 Dog 定义 Eat 方法
// 关键点: 因为要修改 Energy，必须使用指针接收者 (*Dog)
func (d *Dog) Eat() {
	// 1. 打印 "Dog [名字] is eating..."
	fmt.Printf("Dog %s is eating...\n",d.Name)
	// 2. 将 Energy 增加 10
	d.Energy += 10
}

// --- 任务二：接口与多态 ---
/*
要求：
1. 定义 Speaker 接口，包含 Speak() 方法。
2. 定义 Cat 结构体。
3. 让 Dog 和 Cat 都实现 Speak() 方法。
*/
// TODO 4: 定义 Speaker 接口
type Speaker interface {
	// 定义 Speak 方法签名，不需要参数，无返回值（或者按需设计，这里假设无返回值）
	Speak()
}

// TODO 5: 定义 Cat 结构体，包含 Name (string)
type Cat struct {
	// 补充字段
	Name string
}

// TODO 6: 为 Cat 实现 Speak 方法
// 打印 "Cat [名字]: 喵喵喵"
func (c Cat) Speak() {
	// 补充代码
	fmt.Printf("Cat %s：喵喵喵\n",c.Name)
}

// TODO 7: 为 Dog 实现 Speak 方法 (为了满足接口)
// 打印 "Dog [名字]: 汪汪汪"
// 注意：虽然上面有了 Bark，但接口要求的是 Speak，所以需要单独实现
func (d Dog) Speak() {
	// 补充代码
	fmt.Printf("Dog %s：汪汪汪\n",d.Name)
}

// 一个普通的函数，接收接口类型
func MakeItSpeak(s Speaker) {
	s.Speak()
}
func main() {
	fmt.Println("--- 任务一：电子宠物 Dog ---")
	// 初始化一个 Dog
	myDog := Dog{Name: "旺财", Energy: 50}
	fmt.Printf("初始状态: %s, 能量: %d\n", myDog.Name, myDog.Energy)
	// 调用方法
	myDog.Bark()
	myDog.Eat()
	fmt.Printf("进食后状态: %s, 能量: %d\n", myDog.Name, myDog.Energy)
	// 验证：如果 Eat 使用的是 (d Dog) 而不是 (d *Dog)，这里的能量将仍然是 50
	fmt.Println("\n--- 任务二：多态动物园 ---")
	myCat := Cat{Name: "咪咪"}
	// 创建一个 Speaker 切片，存放不同的动物
	// 注意：这里能同时放入 Dog 和 Cat，是因为它们都实现了 Speaker 接口
	animals := []Speaker{
		myDog, // Dog 实现了 Speak
		myCat, // Cat 实现了 Speak
	}
	for _, animal := range animals {
		// 多态调用：虽然代码一样，但不同对象行为不同
		MakeItSpeak(animal)
	}
}

/*
-------------
预期输出:
-------------
--- 任务一：电子宠物 Dog ---
初始状态: 旺财, 能量: 50
Dog 旺财 is barking: 汪汪!
Dog 旺财 is eating...
进食后状态: 旺财, 能量: 60
--- 任务二：多态动物园 ---
Dog 旺财: 汪汪汪
Cat 咪咪: 喵喵喵
*/
