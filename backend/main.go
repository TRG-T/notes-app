package main

import (
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model	
	Username            string
	Password            string
}
var users []User


func main() {
	godotenv.Load(".env")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432", dbHost, dbUser, dbPassword, dbName)
	
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	db.AutoMigrate(&User{})

	if err != nil {
		fmt.Println("Connecting to database failed")
	} else {
		fmt.Println("Connected to database")
	}
	
	db.Table("users").Find(&users)
	for _, user := range users {
		fmt.Println(user.Username, user.Password)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	r.Run("127.0.0.1:3001")
	
}
