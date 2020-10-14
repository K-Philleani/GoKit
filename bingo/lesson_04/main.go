package main

import "fmt"

// 数组 (数组是值类型)
func main() {
	// 初始化方法一
	var testArray [3]int
	var numArray = [3]int{1, 2}
	var cityArray= [3]string{"北京", "深圳", "上海"}
	fmt.Println(testArray)
	fmt.Println(numArray)
	fmt.Println(cityArray)

	// 初始化方法二
	var testArray2 [3]int
	var numArray2 = [...]int{1, 2}
	var cityArray2 = [...]string{"北京", "深圳", "上海"}
	fmt.Println(testArray2)
	fmt.Println(numArray2)
	fmt.Printf("type of numArray2: %T\n", numArray2)
	fmt.Println(cityArray2)
	fmt.Printf("type of cityArray2: %T\n", cityArray2)

	// 初始化方法三
	a := [...]int{1: 1, 3: 5}
	fmt.Println(a)
	fmt.Printf("type of a:%T\n", a)

	// 数组的遍历
	b := [3]string{"北京", "深圳", "上海"}
	// 方法一
	for i := 0; i < 3; i++ {
		fmt.Println(b[i])
	}
	// 方法二
	for i, v := range b {
		fmt.Println(i, v)
	}

	// 多维数组
	multiArray:= [3][2]string{
		{"北京", "上海"},
		{"深圳", "哈尔滨"},
		{"南京", "浙江"},
	}
	fmt.Println(multiArray)

	// 二维数组的遍历
	for _, v1 := range multiArray {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}

	// 多维数组只有第一层可以使用...来让编译器推导数组长度
	c := [...][2]string{
		{"北京", "上海"},
		{"深圳", "哈尔滨"},
		{"南京", "浙江"},
	}
	for _, v1 := range c {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}



}
