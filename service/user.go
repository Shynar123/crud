package service

import (
	"crud/userdb"
	u "crud/userdb"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	CreateUser(c *gin.Context)
	EditUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetAllUsers()
}
type UserServiceImpl struct {
	userRepository u.UserDB
}

func (u UserServiceImpl) CreateUser(c *gin.Context) {
	username := c.Param("username")
	err:=u.userRepository.CreateUserinDB(username)
	if err!=nil{
		c.String(http.StatusBadRequest,err.Error())
	}
	users:=userdb.User{Username: username}
	c.JSON(http.StatusOK, users)
	// username := c.Param("")
	// u.CreateUserinDB(username)
}
func (u UserServiceImpl) EditUser(c *gin.Context) {
	username := c.Param("username")
	err:=u.userRepository.EditUserinDB(username)
	if err!=nil{
		c.String(http.StatusNotFound,err.Error())

	}
	users:=userdb.User{Username: username}
	c.JSON(http.StatusOK, users)
}

func (u UserServiceImpl) DeleteUser(c *gin.Context) {
	username := c.Param("username")
	err:=u.userRepository.DeleteUserinDB(username)
	if err!=nil{
		fmt.Println("error in delete user:",err)
	}
	users:=userdb.User{Username: username}
	c.JSON(http.StatusOK, users)
}

func (u UserServiceImpl) GetAllUsers() []u.User {
	users, err := u.userRepository.GetAllUsersfromDB()
	if err != nil {
		fmt.Println("Get users error:", err)
	}
	// c.JSON(http.StatusOK, users)
	//users?
	return users
}
