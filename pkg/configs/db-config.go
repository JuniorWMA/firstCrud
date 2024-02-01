package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDB() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	d, err := gorm.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		fmt.Println("Error opening DB ")
		log.Fatal(err)
	}
	DB = d
}

func GetConnectDB() *gorm.DB {
	return DB
}
