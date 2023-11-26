package main

import (
	"crud/service"
	"crud/userdb"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	r := setupRouter()

	r.Run(":3000")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("ui/html/*")

	userService := service.NewUsersService(userdb.UserRepositoryInit())
	r.GET("/", func(c *gin.Context) {
		users := userService.GetAllUsers()
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
	r.POST("user/create", func(c *gin.Context) {
		type Req struct {
			Username string `form:"username"`
		}

		var req Req
		err := c.Bind(&req)
		if err != nil {

			c.Status(http.StatusBadRequest)
			return
		}

		// Call the CreateUser method from your UserService
		_, err = userService.CreateUser(req.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
		// c.JSON(http.StatusOK, user)
		c.Redirect(http.StatusMovedPermanently, "/")
	})
	r.POST("user/delete", func(c *gin.Context) {
		type Req struct {
			ID string `form:"id"`
		}

		var req Req
		err := c.Bind(&req)
		if err != nil {

			c.Status(http.StatusBadRequest)
			return
		}

		// Call the CreateUser method from your UserService
		err = userService.DeleteUser(req.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}
		// c.JSON(http.StatusOK, user)
		c.Redirect(http.StatusMovedPermanently, "/")
	})
	r.POST("/user/edit", func(c *gin.Context) {
		type Req struct {
			Username string `form:"username"`
			ID       string `form:"id"`
		}

		var req Req
		err := c.Bind(&req)
		if err != nil {

			c.Status(http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(req.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to convert to int"})
			return
		}
		user := &userdb.User{id, req.Username}
		// Call the CreateUser method from your UserService
		err = userService.EditUser(*user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit user"})
			return
		}
		// c.JSON(http.StatusOK, user)
		c.Redirect(http.StatusMovedPermanently, "/")
	})
	r.GET("/users_json", func(c *gin.Context) {
		users := userdb.User{}
		c.JSON(http.StatusOK, users)
	})

	return r
}
