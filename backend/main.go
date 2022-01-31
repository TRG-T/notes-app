package main

import (
	"fmt"
	"os"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model	
	Username            string
	Password            string
}
var user User

type Note struct {
	gorm.Model
	Title				string
	Description			string
	OwnerID				uint
}
var notes []Note

type LoginRequestBody struct {
	Username string
	Password string
}

type NoteRequestBody struct {
	Title string
	Description string
	User string
}

func registerUser(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var requestBody LoginRequestBody
		err := c.BindJSON(&requestBody)
		if err != nil {
			fmt.Println("Something went wrong")
		}
	
		bytes, _ := bcrypt.GenerateFromPassword([]byte(requestBody.Password), 14)
		checkUser := db.Table("users").Where("username = ?", requestBody.Username).First(&user)
		if checkUser.RowsAffected == 1 {
			c.JSON(200, "User with this name already exist")
		} else {
			db.Table("users").Create(&User { Username: requestBody.Username, Password: string(bytes) })
			c.JSON(200, "Success")
		}
	}
	return gin.HandlerFunc(fn)
}

func loginUser(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var requestBody LoginRequestBody
		err := c.BindJSON(&requestBody)
		if err != nil {
			fmt.Println("Something went wrong")
		}
	
		result := db.Table("users").Where("username = ?", requestBody.Username).First(&user)
		if result.RowsAffected == 0 {
			c.JSON(200, "User not found")
		} else if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password)) != nil {
			c.JSON(200, "Wrong password")
		} else {
			c.JSON(200, "Success")
		}
	}
	return gin.HandlerFunc(fn)
}

func addNote(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var requestBody NoteRequestBody
		err := c.BindJSON(&requestBody)
		if err != nil {
			fmt.Println("Something went wrong")
		}

		db.Table("users").Where("username = ?", requestBody.User).First(&user)
		db.Table("notes").Create(&Note { Title: requestBody.Title, Description: requestBody.Description, OwnerID: user.ID })
		c.JSON(200, "Success")
	}
	return gin.HandlerFunc(fn)
}

func getNotes(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		params := c.Request.URL.Query()

		db.Table("users").Where("username = ?", params["user"]).First(&user)
		db.Table("notes").Where("owner_id = ?", user.ID).Find(&notes)
		c.JSON(200, notes)
	}
	return gin.HandlerFunc(fn)
}


func main() {
	godotenv.Load(".env")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432", dbHost, dbUser, dbPassword, dbName)

	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	db.AutoMigrate(&User{}, &Note{})

	if err != nil {
		fmt.Println("Connecting to database failed")
	} else {
		fmt.Println("Connected to database")
	}

	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/login", loginUser(db))
	r.POST("/register", registerUser(db))
	r.POST("/add-note", addNote(db))
	r.GET("/get-notes", getNotes(db))
	r.Run("127.0.0.1:3001")
	
}
