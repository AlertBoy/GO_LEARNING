package main

import (
	"fmt"
	redis "github.com/garyburd/redigo/redis"
	"time"
	"unsafe"
)

type Db struct {
	Conn redis.Conn
}

func NewRedisClinet(addr string) *Db {
	c, err := redis.Dial("tcp",
		"127.0.0.1:6379",
		redis.DialDatabase(1))
	if err != nil {
		fmt.Println(err)
	}
	return &Db{Conn: c}
}

type SubscribeCallback func(channel, message string)

// 实现了redis的增删改查
// 下载地址go get github.com/garyburd/redigo/redis

func main() {
	c, err := redis.Dial("tcp",
		"127.0.0.1:6379",
		redis.DialDatabase(1))
	if err != nil {
		fmt.Println("Connect to redis failed ,cause by >>>", err)
		return
	}
	defer c.Close()

	// 设置值
	r, err := c.Do("SET", "test-Key", "test-Value")
	if err != nil {
		fmt.Println("redis set value failed >>>", err)
	}
	fmt.Printf("redis seting result: %s", r)

	// 获取值
	exists, err := redis.Bool(c.Do("EXISTS", "test-Key"))
	if err != nil {
		fmt.Println("illegal exception")
	}

	// 设置过期时间EX,5秒
	_, err = c.Do("SET", "test-Key2", "test-Value", "EX", "5")
	if err != nil {
		fmt.Println("redis set value failed >>>", err)
	}

	fmt.Printf("exists or not: %v \n", exists)

	//del kv
	_, err = c.Do("DEL", "test-Key")
	if err != nil {
		fmt.Println("redis delelte value failed >>>", err)

	}

	// 发布订阅
	conn := redis.PubSubConn{c}

	m := make(map[string]SubscribeCallback)

	m["123"] = func(channel, message string) {
		fmt.Println("123123123123213", channel, message)
	}

	err = conn.Subscribe("123456")
	if err != nil {
		fmt.Println("redis Subscribe error.")
	}
	fmt.Println("redis Subscribe 1.")
	go func() {
		for {
			fmt.Println("redis Publish ")
			_, err = c.Do("Publish", "123", "hello")
			if err != nil {
				println(err)
			}
			time.Sleep(2000)
		}
	}()
	fmt.Println("redis Subscribe 2.")
	go func() {
		for {
			fmt.Println("redis Receive 2.")
			switch res := conn.Receive().(type) {
			case redis.Message:
				channel := &res.Channel
				message := (*string)(unsafe.Pointer(&res.Data))
				m[*channel](*channel, *message)
			case redis.Subscription:
				fmt.Printf("%s: %s %d\n", res.Channel, res.Kind, res.Count)
			case error:
				continue
			}
		}
	}()

	fmt.Println("redis Subscribe 3.")
	<-make(chan int)
}
