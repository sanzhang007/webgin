package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	cmd := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(cmd), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	DB.AutoMigrate(&Node1{})
}
