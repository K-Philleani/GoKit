package main

import "fmt"

// 结构体模拟实现继承

// 动物类
type animal struct {
	name string
}
func (a animal) move() {
	fmt.Printf("%s 会动\n", a.name)
}

type dog struct {
	feet uint8
	animal
}
func (d dog) wang() {
	fmt.Printf("%s 会叫\n", d.name)

}

func main() {
	d1 := dog{
		animal:	animal{ name: "tom" },
		feet: 2,
	}
	fmt.Println(d1)
	d1.move()
	d1.wang()

}
