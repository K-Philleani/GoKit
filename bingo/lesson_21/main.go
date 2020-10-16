package main

import "fmt"

// 结构体
type person struct {
	name string
	city string
	age  int
}
func main() {
	var p person
	p.name = "kk"
	p.city = "黑龙江"
	p.age = 23
	fmt.Printf("p=%v\n", p)
	fmt.Printf("p=%#v\n", p)

	// 匿名结构体
	var user struct{
		Name string
		Age int
	}
	user.Name = "jj"
	user.Age = 18
	fmt.Printf("%#v\n", user)

	// 创建指针类型结构体
	var p2 = new(person)
	fmt.Printf("%T\n", p2)
	fmt.Printf("p2=%#v\n", p2)

	// 取结构体的地址实例化
	p3 := &person{}
	fmt.Printf("%T\n", p3)
	fmt.Printf("p3=%#v\n", p3)
	p3.name = "kp"
	p3.age = 20
	p3.city = "哈尔滨"
	fmt.Printf("p3=%#v\n", p3)

	// 使用键值对初始化
	p5 := person{
		name: "晚熟的人",
		city: "哈尔滨",
		age:  23,
	}
	fmt.Printf("p5=%#v\n", p5)
}
