package main

import (
	"fmt"
	"gorm-example/model"
	"gorm-example/storate"
	"gorm.io/gorm"
)

const urlConnect string = "gormMysql://localhost:33061"

func main() {
	dns := "root:123456@tcp(localhost:33061)/testuser?charset=utf8mb4&parseTime=True&loc=Local"
	db := storate.New(dns)

	//Migration
	db.AutoMigrate(&model.User{}, &model.Address{})

	// First create
	findUsers(db)
}

func findUsers(db *gorm.DB)  {
	users := make([] model.User, 0)
	db.Find(&users)
	for _, user:= range users {
		fmt.Printf("%d - %s \n", user.ID, user.Name)
	}
}

func createUser(db *gorm.DB) {
	st := "main street2"
	user := model.User{
		Name: "UserTest2",
		Email: "email2@gmail.com",
		Address: model.Address{
		Street: &st,
	}}
	result := db.Create(&user)

	fmt.Println(user.ID)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
}
