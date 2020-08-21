package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与JSON
// 1.序列化：把Go中的结构体变量 --> json格式的字符串
// 2.反序列化：json的字符串 --> Go能够识别的结构体变量
type person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	p1 := person{
		Name: "tom",
		Age: 24,
	}
	// 序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err: %v", err)
		return
	}
	fmt.Printf("%#v\n",string(b))
	// 反序列化
	str := `{"name": "人生", "age": 20}`
	var p2 person
	json.Unmarshal([]byte(str), &p2)
	fmt.Printf("%#v\n", p2)
}
