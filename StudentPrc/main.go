package main

import (
	"fmt"
	"os"
)
type student struct {
	id int64
	name string
}

var allStudent map[int64]*student

func newStudent(id int64, name string) *student{
	return &student{
		id: id,
		name: name,
	}
}

func findAllStudent() {
	for k, v := range allStudent {
		fmt.Printf("学号：%d 姓名：%s\n", k, v.name)
	}
}
func addStudent() {
	var id int64
	var name string
	fmt.Print("请输入学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名：")
	fmt.Scanln(&name)
	s := newStudent(id, name)
	allStudent[id] = s
}
func deleteStudent() {
	var id int64
	fmt.Print("请输入要你删除的学生学号：")
	fmt.Scanln(&id)
	delete(allStudent, id)
}
func main() {
	allStudent = make(map[int64]*student, 50)

	for {
		fmt.Print("==================学生管理系统==================")
		fmt.Println(`
		1.查看全部学生
		2.添加学生
		3.删除学生
		4.退出系统
		`)
		fmt.Print("请选择你要进行的操作:")
		var index int64
		fmt.Scanln(&index)
		fmt.Println("你得选择是：", index)
		switch index {
		case 1:
			findAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		}
	}
}
