package globalInit

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var RedisClient *redis.Client

func RedisInit() {
	host:=viper.Get("redis.host")
	port:=viper.GetInt("redis.port")
	addr:= fmt.Sprintf("%s:%d",host,port)
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("redis错误：%s",err))
	}
}