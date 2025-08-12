package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

func main() {
	//启一个redis客户端
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()

	//测试连接
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("无法连接到Redis", err)
	}
	fmt.Println("连接成功")

	//string操作 设置和获取字符串， 过期时间10秒
	err = redisClient.Set(ctx, "name", "miao", 10*time.Second).Err()
	if err != nil {
		log.Fatal("set error", err)
	}
	val, err := redisClient.Get(ctx, "name").Result()
	if err != nil {
		log.Fatal("get error", err)
	}
	fmt.Println("name:", val)

	//检查key是否存在
	result, err := redisClient.Exists(ctx, "name").Result()
	if err != nil {
		log.Fatal("exists error", err)
	}
	fmt.Println("exists:", result)

	// hash操作
	user := map[string]interface{}{
		"name":  "zhangsan",
		"age":   18,
		"email": "zhangsan@163.com",
	}
	err = redisClient.HSet(ctx, "user:1001", user).Err()
	if err != nil {
		log.Fatal("hset error", err)
	}
	//获取hash字段
	val, err = redisClient.HGet(ctx, "user:1001", "name").Result()
	if err != nil {
		log.Fatal("hget error", err)
	}
	fmt.Println("name:", val)

	//列表操作
	//从列表左侧插入元素
	err = redisClient.LPush(ctx, "list:1001", "a", "b", "c").Err()
	if err != nil {
		log.Fatal("lpush error", err)
	}
	//从列表左侧弹出元素
	val, err = redisClient.LPop(ctx, "list:1001").Result()
	if err != nil {
		log.Fatal("lpop error", err)
	}
	fmt.Println("val:", val)
	//获取列表所有元素
	values, err := redisClient.LRange(ctx, "list:1001", 0, -1).Result()
	if err != nil {
		log.Fatal("lrange error", err)
	}
	fmt.Println("list values:", values)

	//集合操作
	redisClient.SAdd(ctx, "set:1001", "a", "b", "a")
	//获取集合所有元素
	values, err = redisClient.SMembers(ctx, "set:1001").Result()
	if err != nil {
		log.Fatal("smembers error", err)
	}
	fmt.Println("set values:", values)

	//事务操作
	_, err = redisClient.TxPipelined(ctx, func(p redis.Pipeliner) error {
		p.Set(ctx, "count", 0, 0)
		p.Incr(ctx, "count")
		p.Incr(ctx, "count")
		return nil
	})
	if err != nil {
		log.Fatal("pipelined error", err)
	}
	counter, err := redisClient.Get(ctx, "count").Result()
	if err != nil {
		log.Fatal("get error", err)
	}
	fmt.Println("count:", counter)

	//删除键
	err = redisClient.Del(ctx, "count").Err()
	if err != nil {
		log.Fatal("del error", err)
	}
	fmt.Println("count:", redisClient.Get(ctx, "count"))

	//关闭连接
	redisClient.Close()
}
