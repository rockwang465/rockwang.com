package main

import (
	"context"
	"fmt"
	"time"

	//"github.com/go-redis/redis/v7"  // 昌喜用的这个版本
	redis "github.com/go-redis/redis/v8"
)

func newFailoverClient() {
	var ctx = context.Background()
	// See http://redis.io/topics/sentinel for instructions how to
	// setup Redis Sentinel.
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName: "mymaster",
		//SentinelAddrs: []string{":26379"},
		SentinelAddrs: []string{"10.151.3.114:30379"},
		Password:      "9@mXWb%m8w5n",
	})

	statusCmd := rdb.Ping(ctx)
	fmt.Println("statusCmd:", *statusCmd)
	rdb.Set(ctx, "rock", "good", time.Second*1)
	value := rdb.Get(ctx, "rock").Val()
	fmt.Println("rock value is :", value)
}

func main() {
	newFailoverClient()
}

// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build test_redis_sentinel_connect.go
// scp test_redis_sentinel_connect root@10.151.3.114:/root/
// ./test_redis_sentinel_connect  // 结果如下:
//redis: 2021/01/04 12:51:24 sentinel.go:610: sentinel: new master="mymaster" addr="10.244.0.48:6379"
//statusCmd: {{0xc000194000 [ping] <nil> 0 <nil>} PONG}
//rock value is : good
