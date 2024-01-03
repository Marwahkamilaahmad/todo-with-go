package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	const mySql = "root:@tcp(localhost:3306)/go_fiber_one"
	DSN := mySql

	database, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic("can't connect to database")
	}
	
	DB = database

	fmt.Println("connected to database")
}
