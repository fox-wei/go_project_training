package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	//*创建client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	//*set key-value
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	//*get key
	val, err := client.Get("key").Result()
	if err != nil {
		panic(val)
	}

	fmt.Printf("key:%v\n", val)

	//*close
	err = client.Close()
	if err != nil {
		panic(err)
	}
}
