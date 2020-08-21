package main

import (
	"fmt"
	"os"
)

type student struct {
	id int64
	name string
}

type Manager struct {
	allStudent map[int64]student
}

func (m Manager) showStudent() {
	for _, v := range m.allStudent {
		fmt.Printf("学号：%d 姓名：%s\n", v.id, v.name)
	}
}
func (m Manager) addStudent() {
	var (
		id int64
		name string
	)
	fmt.Print("请输入学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)

	newStu := student{
		id: id,
		name: name,
	}
	m.allStudent[id] = newStu
}
func (m Manager) deleteStudent() {
	var id int64
	fmt.Print("请输入要删除的学生学号：")
	fmt.Scanln(&id)
	fmt.Printf("删除的学生是：%s\n", m.allStudent[id].name)
	delete(m.allStudent, id)
}
func (m Manager) modifyStudent() {
	var id int64
	fmt.Print("请输入要修改的学生的学号：")
	fmt.Scanln(&id)

	obj, ok := m.allStudent[id]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	fmt.Printf("你要修改的学生————学号：%d 姓名： %s\n", obj.id, obj.name)
	fmt.Printf("请输入学生的新名字：")
	var newName string
	fmt.Scanln(&newName)
	obj.name = newName
	m.allStudent[id] = obj
}

var smr Manager

func menu() {
	fmt.Println("===欢迎使用学生管理系统===")

	fmt.Println("1.查看所有学生信息")
	fmt.Println("2.添加新的学生信息")
	fmt.Println("3.删除学生信息")
	fmt.Println("4.修改学生信息")
	fmt.Println("5.退出系统")
}

func main() {
	smr = Manager{
		allStudent: make(map[int64]student, 100),
	}
	for {
		var index int64
		menu()
		fmt.Print("请输入你要进行的操作的序号:")
		fmt.Scanln(&index)
		fmt.Printf("你输入的序号是：%d\n", index)
		switch index {
		case 1:
			smr.showStudent()
		case 2:
			smr.addStudent()
		case 3:
			smr.deleteStudent()
		case 4:
			smr.modifyStudent()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("错误")
		}
	}
}
