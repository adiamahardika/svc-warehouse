package main

import (
	"fmt"
	"svc-warehouse/router"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:password@tcp(127.0.0.1:3306)/warehouse_guntur?charset=utf8mb4&parseTime=True&loc=Local"
	db, error := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if error != nil {
		fmt.Println("Connection to db has been error!")
	} else {
		fmt.Println("Connection to db succeed!")
	}

	router.AllRouter(db)
}
