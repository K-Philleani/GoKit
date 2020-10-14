package main

import "fmt"

// 指针
func main() {
	// 指针地址和指针类型
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a)
	fmt.Printf("b:%p type: %T\n", b, b)
	fmt.Println(&b)

	// 指针取值
	c := *b
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)

	// new
	d := new(int)
	e := new(bool)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Println(*d)
	fmt.Println(*e)

	// make

	f := make(map[string]int)
	f["p1"] = 10
	fmt.Println(f)
}

