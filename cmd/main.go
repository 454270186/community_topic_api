package main

import (
	"github.com/454270186/CommuTopicPage/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var dbConfig config.Config
var db *gorm.DB
var rdb *redis.Client

func main() {
	r := NewRouter()

	r.Run()
}