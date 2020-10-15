package main

import (
	"fmt"
	"strconv"
)

// strconv包: 实现了基本数据类型与其字符串表示的转换
func main() {
	// Atoi()函数用于将字符串类型的整数转换为int类型
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type: %T value:%#v\n", i1, i1)
	}

	// Itoa()函数用于将int类型数据转换为对应的字符串表示
	i2 := 100
	s2 := strconv.Itoa(i2)
	fmt.Printf("type:%T value:%#v\n", s2, s2)

	// Parse系列函数
	b, err := strconv.ParseBool("true")
	fmt.Println(b)
	f, err := strconv.ParseFloat("3.1415", 64)
	fmt.Println(f)
	i, err := strconv.ParseInt("-2", 10, 64)
	fmt.Println(i)
	u, err := strconv.ParseUint("2", 10, 64)
	fmt.Println(u)
}
