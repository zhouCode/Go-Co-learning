package main

import "fmt"

// ==========================================
// 1. 基础结构体与“组合” (Composition)
// ==========================================

// BaseCharacter 基础角色，包含通用的属性
// Go 没有 class，也没有 extends，我们用 struct 来存数据
type BaseCharacter struct {
	Name  string
	HP    int
	MaxHP int
}

// Warrior 战士
type Warrior struct {
	// 【关键点】：匿名嵌入 (Anonymous Embedding)
	// 我们直接把 BaseCharacter 放在这里，Warrior 就自动“拥有”了它的字段
	// 这就是 Go 实现“继承”的方式：组合优于继承
	BaseCharacter

	AttackPower int // 战士特有的：攻击力
}

// Mage 法师
type Mage struct {
	BaseCharacter

	Mana    int // 法师特有的：魔法值
	MaxMana int
}

// ==========================================
// 2. 方法 (Methods) - 定义角色的行为
// ==========================================

// 为 BaseCharacter 定义一个通用方法
// 接收者 (b *BaseCharacter) 是指针，因为受伤会修改 HP
func (b *BaseCharacter) TakeDamage(dmg int) {
	b.HP -= dmg
	if b.HP < 0 {
		b.HP = 0
	}
	fmt.Printf(" -> [%s] 受到 %d 点伤害，剩余 HP: %d/%d\n", b.Name, dmg, b.HP, b.MaxHP)
}

// Warrior 的攻击方法
func (w *Warrior) Attack(target *BaseCharacter) {
	fmt.Printf("⚔️ 战士 [%s] 挥舞大剑！\n", w.Name)
	// 战士造成物理伤害
	target.TakeDamage(w.AttackPower)
}

// Mage 的攻击方法
func (m *Mage) Attack(target *BaseCharacter) {
	if m.Mana < 10 {
		fmt.Printf("💨 法师 [%s] 想要施法，但是没蓝了...\n", m.Name)
		return
	}

	m.Mana -= 10
	fmt.Printf("🔥 法师 [%s] 发射火球术！(消耗10点蓝，剩余: %d)\n", m.Name, m.Mana)
	// 法师造成魔法伤害 (固定 50 点)
	target.TakeDamage(50)
}

// ==========================================
// 3. 接口 (Interface) - 多态的实现
// ==========================================

// Attacker 接口
// 任何拥有 Attack(target *BaseCharacter) 方法的类型，都是 Attacker
type Attacker interface {
	Attack(target *BaseCharacter)
}

func main() {
	// --- 1. 初始化对象 ---
	// 也就是“实例化”，但在 Go 里只是创建结构体变量

	arthur := &Warrior{
		BaseCharacter: BaseCharacter{Name: "亚瑟", HP: 100, MaxHP: 100},
		AttackPower:   30,
	}

	angela := &Mage{
		BaseCharacter: BaseCharacter{Name: "安琪拉", HP: 60, MaxHP: 60},
		Mana:          100,
		MaxMana:       100,
	}

	// 定义一个可怜的靶子
	dummy := &BaseCharacter{Name: "木桩", HP: 500, MaxHP: 500}

	// --- 2. 使用嵌入特性 ---
	// 注意：Warrior 并没有直接定义 Name，但是我们可以直接访问 arthur.Name
	// Go 自动帮我们转发到了 arthur.BaseCharacter.Name
	fmt.Println("战斗开始！玩家:", arthur.Name, "vs", angela.Name)

	// --- 3. 接口与多态 ---
	// 创建一个 Attacker 切片，把战士和法师都放进去
	// 因为他们都实现了 Attack 方法，所以他们都是 "Attacker"
	party := []Attacker{arthur, angela}

	fmt.Println("\n--- 全员集火攻击木桩 ---")
	for _, hero := range party {
		// 这里的 hero 是接口类型，Go 会自动调用对应的具体实现
		hero.Attack(dummy)
	}

	fmt.Println("\n--- 第二轮攻击 ---")
	// 再打一次，看看法师扣蓝
	angela.Attack(dummy)
}
