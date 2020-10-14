package main

import "fmt"

// 切片
func main() {
	// 声明切片类型
	var a []string
	var b = []int{}
	var c = []bool{false, true}
	var d = []bool{false, true}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(a == nil)
	fmt.Println(b == nil)
	fmt.Println(c == nil)

	// 切片的长度和容量
	 s := [5]int{1, 2, 3, 4, 5}
	 x := s[1:3]
	 fmt.Printf("s:%v len(s):%v cap(s):%v\n", x, len(x), cap(x))

	 // 使用make()函数构造切片
	 y := make([]int, 2, 10)
	 fmt.Println(y)
	 fmt.Println(len(y))
	 fmt.Println(cap(y))

	 // 判断切片是否为空(切片之间是不能比较的，我们不能使用==操作符来判断两个切片是否含有全部相等元素)
	 // nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0。但是我们不能说一个长度和容量都是0的切片一定是nil
	 // 切片唯一合法的比较操作是和nil比较(要判断一个切片是否是空的，要是用len(s) == 0来判断，不应该使用s == nil来判断)
	 var s1 []int
	 s2 := []int{}
	 s3 := make([]int, 0)
	 fmt.Println(s1 == nil, s1)
	 fmt.Println(s2 == nil, s2)
	 fmt.Println(s3 == nil, s3)

	 // append()方法为切片添加元素
	 s4 := append(s1, 1)
	 fmt.Println(s4)
	 s5 := append(s2, 2, 3, 4)
	 fmt.Println(s5)

	 // 使用copy()函数复制切片
	 s6 := make([]int, 3)
	 copy(s6, s5)
	 fmt.Println(s6)

	 // 从切片中删除元素
	 s6 = append(s6[0:1], s6[2:]...)
	 fmt.Print(s6)


}
