package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Retrieve all *:Guild:* keys
	keyRes, _ := rdb.Keys(ctx, "player-id-1:Guild:*").Result()
	fmt.Printf(">> guild for 'player-id-1': %s\n", keyRes)

	// Retrieve permissions for each guild
	// TODO: use MULTI to retrieve all at once
	for _, guildKey := range keyRes {
		fmt.Println("== " + guildKey + "==")
		guildRes, _ := rdb.HGetAll(ctx, guildKey).Result()
		fmt.Println(guildRes)
	}
}
