package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)
// 文件读取操作

// 打开文件，Read读取
func fileRead() {
	// 只读方式打开当前目录下的file.txt文件
	file, err := os.Open("./file.txt")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	// 关闭文件
	defer file.Close()
	// 使用Read方法读取数据
	var tmp = make([]byte, 128)
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件读取完了")
		return
	}
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Printf("读取了%d字节数据\n", n)
	fmt.Println(string(tmp[:n]))
}

// 循环读取
func cirRead() {
	file, err := os.Open("./read.txt")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	defer file.Close()
	// 循环读取文件
	var content []byte
	var tmp = make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读取完了")
			break
		}
		if err != nil {
			fmt.Println("read file filed, err:", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}

// bufio读取文件
func bufioRead() {
	file, err := os.Open("./read.txt")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Println(line)
	}
}

// ioutil读取整个文件
func ioutilRead() {
	content, err := ioutil.ReadFile("./file.txt")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}

func main() {
	//fileRead()
	//cirRead()
	//bufioRead()
	ioutilRead()
}
