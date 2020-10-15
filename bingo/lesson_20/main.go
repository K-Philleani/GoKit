package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

// Redis
var rdb *redis.Client

// 初始化连接
func initClient() (err error){
	rdb = redis.NewClient(&redis.Options{
		Addr: "124.70.71.78:6379",
		Password: "",
		DB: 0,
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func redisDemo() {
	err := rdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}
	val, err := rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score", val)

	val2, err := rdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}
}

func main() {
	err := initClient()
	if err != nil {
		panic(err)
	}
	log.Println("redis success")
	redisDemo()
}
