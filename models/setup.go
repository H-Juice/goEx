package models //import "github.com/JenYata/goEx"

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	dsn := "user:password@tcp(localhost:3306)/database?parseTime=True"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	database.AutoMigrate(&Book{}) //Bind Book table in Database

	DB = database
}
