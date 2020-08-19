// 一个简易的学生管理系统(函数版：查看，新增，删除学生)
package main

import (
	"fmt"
	"os"
)

type student struct {
	id int64
	name string
}

var (
	allStudent map[int64]*student
)

func newStudent(id int64, name string) *student{
	return &student{
		id: id,
		name: name,
	}
}

func showAllStudent() {
	// 打印所有学生
	for k, v := range allStudent{
		fmt.Printf("学号：%d 姓名：%s\n", k, v.name)
	}
}

func addStudent() {
	// 添加一个新的学生
	// 1.创建一个学生
	var (
		id int64
		name string
	)
	fmt.Print("请输入学生的学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入学生的姓名：")
	fmt.Scanln(&name)
	s := newStudent(id, name)
	// 2.添加到allStudent中
	allStudent[id] = s

}
func deleteStudent() {
	// 1.输入要删除的学生的学号
	var (
		deleteId int64
	)
	fmt.Print("请输入学号是哪个学号：")
	fmt.Scanln(&deleteId)
	// 2.重allStudent中去除
	delete(allStudent, deleteId)
}
func main() {
	allStudent = make(map[int64]*student, 50) // 初始化
	for {
		// 1.打印菜单
		fmt.Println("欢迎光临学生管理系统！")
		fmt.Println(`
		1.查看所用学生 
		2.新增学生
		3.删除学生
		4.退出`)
		fmt.Print("请选择你要进行的操作:")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项\n", choice )
		// 2.等待用户选择
		// 3.执行用户选择的操作
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)// 退出
		default:
			fmt.Println("发生错误")
		}
	}

}
