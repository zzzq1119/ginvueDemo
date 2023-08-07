package common

import (
	"fmt"
	"ginvueDome/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	host := "localhost"
	port := "3306"
	database := "dome"
	username := "root"
	password := "748596123"
	charset := "utf8mb4"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	db, err := gorm.Open(mysql.Open(args))

	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&model.User{})

	return db
}
