package connection

import (
	"github.com/gookit/config/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlConnect() *gorm.DB {

	//host := config.String("db_host")
	port := config.String("db_port")
	username := config.String("db_username")
	password := config.String("db_password")
	database := config.String("db_name")

	dsn := username + ":" + password + "@tcp(127.0.0.1:" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect db")
	}

	return db

}
