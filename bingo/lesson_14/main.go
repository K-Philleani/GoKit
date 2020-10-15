package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 文件写入操作
/**
	os.O_CREATE 创建文件
	os.O_WRONLY 只写
	os.O_RDONLY 只读
	os.O_RDWR 读写
	os.O_TRUNC 清空
	os.O_APPEND 追加
	perm：文件权限，一个八进制数。r（读）04，w（写）02，x（执行）01。
 */

// Write和WriteString
func WriteFile() {
	file, err := os.OpenFile("os.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	str := "Hello, Golang!"
	file.Write([]byte(str))
	file.WriteString("\nHello, File!!")
}

// bufio.NewWriter
func bufioFile() {
	file, err := os.OpenFile("buf.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("Hello, Golang\n") //将数据先写入缓存
	}
	writer.Flush() //将缓存中的内容写入文件
}

// ioutil.WriteFile
func ioutilFile() {
	str := "Hello. Golang"
	err := ioutil.WriteFile("./io.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

func main() {
	//WriteFile()
	//bufioFile()
	ioutilFile()
}
