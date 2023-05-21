package main

import (
	"fmt"

	"github.com/454270186/CommuTopicPage/config"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConfig config.Config
var db *gorm.DB

func init() {
	v := viper.New()
	v.SetConfigFile("./config.json")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&dbConfig); err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", 
					   dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.DBname, dbConfig.Port)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func main() {
	r := NewRouter()

	r.Run()
}