package main

import (
	"fmt"
	"time"
)

// time

// 时间类型
func timeDemo() {
	now := time.Now()
	fmt.Printf("current time:%v\n", now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
// 时间戳
func timeStampDemo() {
	now := time.Now()
	timeStamp1 := now.Unix() //时间戳
	timeStamp2 := now.UnixNano() // 纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timeStamp1)
	fmt.Printf("current timestamp2:%v\n", timeStamp2)
}

// 时间操作
func add() {
	now := time.Now()
	later := now.Add(time.Hour) // 当前时间加1小时后的时间
	fmt.Println(later)
}
// 定时器
func tickDemo() {
	ticker := time.Tick(time.Second) // 定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println(i)
	}
}

// 时间格式化
func formatDemo() {
	now := time.Now()

	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	fmt.Println(now.Format("2006-01-01 03:04:05.000"))
}

// 解析字符串格式的时间
func timeParse() {
	now := time.Now()
	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2020/10/15 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))
}

func main() {
	timeDemo()
	timeStampDemo()
	add()
	//tickDemo()
	formatDemo()
	timeParse()
}
