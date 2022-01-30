package main

import (
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/gin-contrib/cors"
)

type User struct {
	gorm.Model	
	Username            string
	Password            string
}
var users []User

type LoginRequestBody struct {
	Username string
	Password string
}


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
	r.Use(cors.Default())
	r.POST("/login", func(c *gin.Context) {
		var requestBody LoginRequestBody
		err := c.BindJSON(&requestBody)
		if err != nil {
			fmt.Println("Something went wrong")
		}

		fmt.Println(requestBody.Username, requestBody.Password)
	})
	r.Run("127.0.0.1:3001")
	
}
