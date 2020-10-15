package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// map
func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of scoreMap:%T\n", scoreMap)

	userInfo := map[string]string{
		"username": "kk",
		"password": "123456",
	}
	fmt.Println(userInfo)

	// 判断某个键是否存在
	v, ok := userInfo["username"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("信息不存在")
	}

	// map的遍历
	for k, v := range userInfo {
		fmt.Println(k, v)
	}

	// 使用delete()函数删除键值对
	delete(scoreMap, "小明")
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	// 按照指定顺序遍历map
	rand.Seed(time.Now().UnixNano())
	var mp = make(map[string]int, 20)
	for i := 0; i < 20; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		mp[key] = value
	}
	var keys = make([]string, 20)
	for key := range mp {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, mp[key])
	}

	// 元素为map类型的切片
	var mapSlice = make([]map[string]string, 3)
	for k, v := range mapSlice {
		fmt.Printf("index:%d value:%s\n", k, v)
	}
	fmt.Println("defer init")
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "kkk"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "哈哈哈"
	for k, v := range mapSlice {
		fmt.Println(k, v)
	}

	// 值为切片类型的map
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "city"
	v1, ok := sliceMap[key]
	if !ok {
		v1 = make([]string, 0, 2)
	}
	v1 = append(v1, "北京", "上海")
	sliceMap[key] = v1
	fmt.Println(sliceMap)
}
