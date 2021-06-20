package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"unsafe"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	v := "0123456789"
	//v = v+v+v+v+v //50
	//v = v + v  //100
	//v = v + v // 200
	//v = v+v+v+v+v //1k
	//v = v+v+v+v+v //5k
	fmt.Println(unsafe.Sizeof(v))
	fmt.Println(len([]byte(v)))
	fmt.Println(v)
	for i := 0; i < 10000; i++ {
		k := fmt.Sprintf("key%d", i)
		err := rdb.Set(ctx, k, v, 0).Err()
		if err != nil {
			log.Fatalln(err)
		}

	}
	//err := rdb.Set(ctx, "key", "value", 0)
	//if err != nil {
	//  log.Fatalln(err)
	//}
}
