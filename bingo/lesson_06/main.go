package main

import (
	"fmt"
	"log"
	"os"
)

// log


//const (
//	Ldate = 1 << iota
//	Ltime
//	Lmicroseconds
//	Llongfile
//	Lshortfile
//	LUTC
//	LstdFlags = Ldate | Ltime
//)

func main() {
	log.Println("这是一条很普通的日志")
	v := "很普通的"
	log.Printf("这是一条%s日志", v)
	//log.Fatalln("这是一条会触发fatal的日志")
	//log.Panicln("这是一条会触发panic的日志")

	// 配置logger
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	// 配置日志前缀
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("[Kphilleani]")
	log.Println("这是一条很普通的日志。")

	// 配置日志输出位置(通常会把下面的配置操作写到init函数中。)
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[Kphilleani]")
	log.Println("这是一条很普通的日志。")

	// 创建logger
	logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Println("这是自定义的logger记录的日志")

}
