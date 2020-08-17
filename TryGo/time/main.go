package main

import (
	"fmt"
	"time"
)

// time包： 处理时间和日期
func main() {
	fmt.Println("UNIX时间：", time.Now().Unix())
	fmt.Println("time.Now()的时间：", time.Now())
	t := time.Now()
	fmt.Println("将时间：",t, "格式化为RFC3339格式：",t.Format(time.RFC3339))

	time.Sleep(time.Second)
	t1 := time.Now()
	fmt.Println("Time difference：", t1.Sub(t))
}
