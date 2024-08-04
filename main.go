package main

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func redisConn() (rdb *redis.Client) {
	rdb = redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
		Password: "",
		DB:		  1,  // use plant database
	})
	
	return rdb
}

func main() {
	rdb := redisConn()
	ctx := context.Background()
	err := rdb.Publish(ctx, "plants", "Hi, I am a plant!").Err()
	if err != nil {
		panic(err)
	}
	rdb.Close()
}




