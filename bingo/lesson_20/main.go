package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

// redis
var rdb *redis.Client

// 普通连接
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

// 连接Redis哨兵模式
func initFailoverClient() (err error){
	rdb = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName: "master",
		SentinelAddrs: []string{"124.70.71.79:6379"},
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return
	}
	return nil
}

// 连接Redis集群
func initClusterClient() (err error) {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":9090"},
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return
	}
	return nil
}

// set/get
func redisSet() {
	err := rdb.Set("name", "Kphilleani",0).Err()
	if err != nil {
		log.Printf("set failed, err: %v\n", err)
		return
	}
	log.Println("set success")
}
func redisGet() {
	val, err := rdb.Get("name").Result()
	if err != nil {
		fmt.Printf("get failed, err: %v\n", err)
		return
	}
	fmt.Println("name", val)
}

func main() {
	initClient()
	//redisSet()
	redisGet()
}