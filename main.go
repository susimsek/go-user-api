package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"user-api/config"
	"user-api/model"
	"user-api/route"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		log.Fatalln(err)
		fmt.Println("Status:", err)
	}
	defer config.DB.Close()

	config.DB.AutoMigrate(&model.User{})
	r := route.SetupRouter()
	//running
	r.Run()
}
