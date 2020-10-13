package main

import (
	"fmt"
	"math"
	"strings"
)
/**
	基本数据类型
 */

func main() {
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 1010 占位符%b表示二进制

	// 八进制 以0开头
	var b int =077
	fmt.Printf("%o \n", b) // 77

	// 十六进制 以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF

	// 浮点数
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	// 复数
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)

	// 布尔值
	bool1 := true
	var bool2 bool
	fmt.Println(bool1)
	fmt.Println(bool2)

	// 字符串
	// 多行
	s1 := `第一行
第二行
第三行
    `
	fmt.Println(s1)
	// 字符串常用操作
	str := "abcdefg"
	fmt.Println(str)
	// 求长度
	length := len(str)
	fmt.Println(length)
	// 拼接字符串
	str2 := "QWER"
	str3 := str + str2
	fmt.Println(str3)
	fmt.Println(fmt.Sprintf("%s%s", str3, str2))
	// 分割
	var strArr []string = strings.Split(str3, "")
	fmt.Println(strArr)
	// 是否包含
	isContain := strings.Contains(str3, "QWER")
	fmt.Println(isContain)


}
