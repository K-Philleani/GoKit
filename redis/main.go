package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

var client *redis.Client

// 向key的hash中添加元素
func HashSet(key, filed string, data interface{}) {
	value, err := client.HSet(key, filed, data).Result()
	if err != nil {
		log.Println("Redis HSet Error:", err)
		return
	}
	if value {
		log.Println("set success")
	} else {
		log.Println("set failed")
	}
}
// 读取key的hash中元素
func HashGet(key, filed string) {
	value, err := client.HGet(key, filed).Result()
	if err != nil {
		log.Println("Redis HGet Error:", err)
		return
	}
	fmt.Println("value:", value)
}

func ZSet(key string, zlist []redis.Z) {
	num, err := client.ZAdd(key, zlist...).Result()
	if err != nil {
		log.Printf("zadd failed, err: %v\n", err)
		return
	}
	fmt.Println("zadd success:", num)

}

func ZGet(key string, op redis.ZRangeBy) {
	ret, err := client.ZRangeByScoreWithScores(key, op).Result()
	if err != nil {
		fmt.Printf("zrangewithscore failed, err: %v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

func main() {
	client = redis.NewClient(&redis.Options{
		Addr: "124.70.71.78:6379",
		Password: "",
		DB: 0,
	})
	pong, err := client.Ping().Result()
	if err == redis.Nil  {
		log.Println("redis异常")
		return
	} else if err != nil {
		log.Println("redis失败", err)
		return
	} else {
		log.Println(pong)
	}
	//HashSet("cloth", "color", "red")
	//HashGet("cloth", "color")
	//ZSet("zseter", []redis.Z{
	//	redis.Z{Score: 90.0, Member: "Golang"},
	//})
	ZGet("zseter", redis.ZRangeBy{Min: "80", Max: "100"})
}
