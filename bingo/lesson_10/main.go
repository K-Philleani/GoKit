package main

import (
	"errors"
	"fmt"
	"log"
)

// 函数进阶
type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// 高阶函数
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// 函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

func main() {
	var c calculation
	c = add
	fmt.Println(c(1, 2))

	s := sub
	fmt.Println(s(10, 2))

	ret := calc(10, 20, add)
	fmt.Println(ret)
	get, err := do("+")
	if err != nil {
		log.Println("err:", err)
	}
	res := get(5, 2)
	fmt.Println(res)

	// 匿名函数
	anFunc := func(x, y int) {
		fmt.Println(x, y)
	}
	anFunc(1, 2)

	// 自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(1, 10)
}
