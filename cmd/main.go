package main

import (
	"gorm-example/storate"
)

const urlConnect string = "gormMysql://localhost:33061"

func main() {
	dns := "user_root:123456@tcp(localhost:33061)/?charset=utf8mb4&parseTime=True&loc=Local"
	storate.NewDB(dns)
}
