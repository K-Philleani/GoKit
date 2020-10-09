package main

import "fmt"

// 变量和常量
/*
	1. 函数外的每个语句都必须以关键字开始（var、const、func等）
	2. :=不能使用在函数外。
	3. _多用于占位，表示忽略值。
 */

var name string // 仅定义变量，为赋值，默认零值
var age int // 仅定义变量，为赋值，默认零值
var isOK bool // 仅定义变量，为赋值，默认零值

var ( // 仅定义变量，为赋值，默认零值
	a string
	b int
	c bool
	d float32
)

// 常量(恒定不变的值，多用于定义程序运行期间不会改变的那些值)
const pi = 3.1415
const e = 2.7182

const (
	a1 = 1
	b1 = 2
)
// const同时声明多个常量时，如果省略了值则表示和上面一行的值相同
const (
	a2 = 100
	b2
	c2
)

// iota(常量计数器，只能在常量的表达式中使用)
const (
	n1 = iota // 0
	n2
	n3
	n4
)

func main() {
	fmt.Println("name:", name)
	name = "K1"
	fmt.Println("name:", name)
	fmt.Println("age:", age)
	age = 20
	fmt.Println("age:", age)
	fmt.Println("isOk:", isOK)
	isOK = true
	fmt.Println("isOk:", isOK)

	// 短变量方式，声明同时赋值(仅能在函数内使用)
	userName := "k2"
	fmt.Println("userName:", userName)
	// 匿名变量(用于抛弃某个接受到的值, 匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明)
	x, _ := foo()
	_, y := foo()
	fmt.Println("x:", x)
	fmt.Println("y:", y)

	fmt.Println("pi:", pi)
	fmt.Println("e:", e)
	fmt.Println("a1:", a1)
	fmt.Println("a1:", b1)
	fmt.Println("a2:", a2)
	fmt.Println("b2:", b2)
	fmt.Println("c2:", c2)

	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)
	fmt.Println("n4:", n4)

}


func foo() (int, string) {
	return 10, "k3"
}
