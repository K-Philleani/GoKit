package main

import "fmt"

// 函数
func sayHello() {
	fmt.Println("Hello")
}

// 参数
func intSum(x, y int) int {
	return x + y
}
// 可变参数
func intSum2(x ...int) int {
	fmt.Println(x) // x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}
// 返回值
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}
// 返回值命名
func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func main() {
	// 函数的调用
	sayHello()
	sum := intSum(1, 2)
	fmt.Println(sum)
	res := intSum2(1, 2, 3, 4)
	fmt.Println(res)
	res1, res2 := calc(10, 5)
	fmt.Println(res1, res2)
	res3, res4 := calc2(20, 10)
	fmt.Println(res3, res4)
}
