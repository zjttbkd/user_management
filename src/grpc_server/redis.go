package main

import (
	"encoding/json"
	log "github.com/cihub/seelog"
	"github.com/go-redis/redis"
	"time"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "172.18.0.23:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 2000,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func cacheUserInfo(ui *userInfo) {
	value, err := json.Marshal(ui)
	if err != nil {
		log.Error(err)
		return
	}

	_, err = client.SetNX(ui.username, value, 360*time.Second).Result()
	if err != nil {
		log.Error(err)
	}

}

func fetchUserInfo(username string) (*userInfo, error) {
	value, err := client.Get(username).Result()
	if err == redis.Nil {
		return nil, err
	} else if err != nil {
		log.Error(err)
		return nil, err
	}

	var ui userInfo
	err = json.Unmarshal([]byte(value), &ui)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &ui, nil
}

func delUserInfo(username string) {
	_, err := client.Del(username).Result()
	if err != nil {
		log.Error(err)
	}
}
