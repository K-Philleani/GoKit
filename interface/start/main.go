package main

import "fmt"

type Speaker interface {
	speak()
}

type Cat struct {}
type Dog struct {}
type Person struct {}

func (c Cat) speak() {
	fmt.Println("汪汪汪")
}

func (d Dog) speak() {
	fmt.Println("喵喵喵")
}

func (p Person) speak() {
	fmt.Println("Fuck")
}

func da(s Speaker) {
	s.speak()
}


func main() {
	c1 := Cat{}
	d1 := Dog{}
	p1 := Person{}

	da(c1)
	da(d1)
	da(p1)

}
