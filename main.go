package main

import (
	"crud/service"
	"crud/userdb"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db map[string]any

func init() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("db:", err)
	}
}
func main() {

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("ui/html/*")
	userService := &service.UserServiceImpl{}
	r.GET("/", func(c *gin.Context) {
		users := userService.GetAllUsers
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title": "Home Page",
				"users": users,
			},
		)
		// c.ShouldBindWith(binding.Binding)

	})
	r.POST("user/create",func(c *gin.Context){
		// if err := c.ShouldBindJSON(&userForm); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }

		// Call the CreateUser method from your UserService
		user, err := userService.CreateUser()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		// Return the created user as JSON response
		c.JSON(http.StatusOK, user)
	})
	r.POST("/user/edit", service.UserService.CreateUser)
	r.GET("/users_json", func(c *gin.Context) {
		// Respond with JSON
		users := userdb.User{}
		c.JSON(http.StatusOK, users)
	})
	// Get user value
	
	return r
}
