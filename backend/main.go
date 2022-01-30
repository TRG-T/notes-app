package main

import (
	"fmt"
	"os"
	"github.com/gin-contrib/cors"
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
var user User

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

	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/login", func(c *gin.Context) {
		var requestBody LoginRequestBody
		err := c.BindJSON(&requestBody)
		if err != nil {
			fmt.Println("Something went wrong")
		}

		fmt.Println(requestBody.Username, requestBody.Password)
		db.Table("users").Where("username = ?", requestBody.Username).First(&user)
		if requestBody.Password != user.Password {
			c.JSON(200, "Wrong password")
		} else {
			c.JSON(200, "Success")
		}
	})
	r.Run("127.0.0.1:3001")
	
}
