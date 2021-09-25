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
	err := db.AutoMigrate(&model.User{}, &model.Address{})
	if err != nil {
		panic(err)
	}

	// Create User
	//createUser(db)

	// List of users
	//findUsers(db)

	// User By id
	//findUserById(db)

	// Update select fields
	updateById(db)
}

func findUsers(db *gorm.DB) {
	fmt.Printf("Find users  \n")
	users := make([]model.User, 0)
	db.Find(&users)
	for _, user := range users {
		fmt.Printf("%d - %s \n", user.ID, user.Name)
	}
}

func findUserById(db *gorm.DB) {
	fmt.Printf("Find user By id  \n")
	user := model.User{}
	db.First(&user, 3)
	fmt.Printf("%d - %s \n", user.ID, user.Name)
}

func createUser(db *gorm.DB) {
	st := "main street5"
	user := model.User{
		Name:  "UserTest5",
		Email: "email5@gmail.com",
		Address: model.Address{
			Street: &st,
		}}
	result := db.Create(&user)

	fmt.Println(user.ID)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
}

func updateById(db *gorm.DB) {
	user := model.User{}
	user.ID = 4

	db.Model(&user).Updates(
		model.User{Name: "User number 4"})
}
