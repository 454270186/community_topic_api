package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	v := viper.New()
	v.SetConfigFile("./config.json")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&dbConfig); err != nil {
		panic(err)
	}

	// connect to Postgres
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		dbConfig.Host, dbConfig.User, dbConfig.PPassword, dbConfig.DBname, dbConfig.Port)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// connect to Redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     dbConfig.Addr,
		Password: dbConfig.RPassword,
		DB:       dbConfig.DB,
	})
	if err := rdb.FlushAll(context.Background()).Err(); err != nil {
		panic(err)
	}
}