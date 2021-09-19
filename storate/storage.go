package storate

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type dataBase struct {
	db *gorm.DB
}

func New(dns string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Connect to db")
	return db
}


