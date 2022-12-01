package main

import (
	"fmt"

	"belajar-golang/latihan-2/gorm-gin-1/Config"
	"belajar-golang/latihan-2/gorm-gin-1/Models"
	"belajar-golang/latihan-2/gorm-gin-1/Routers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/belajar_gorm_gin_2?charset=utf8&parseTime=True&loc=Local"
	Config.DB, err = gorm.Open("mysql", dsn)

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Todos{})

	router := Routers.SetupRouter()
	router.Run()
}
