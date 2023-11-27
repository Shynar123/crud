package service

import (
	"crud/userdb"
	u "crud/userdb"
	"fmt"
	"strconv"
)

type UserService interface {
	CreateUser(string)
	EditUser(string)
	DeleteUser(string)
	GetAllUsers()
}
type UserServiceImpl struct {
	userRepository u.UserDB
}

func NewUsersService(db u.UserDB) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: db,
	}
}

func (u UserServiceImpl) CreateUser(username string) (u.User, error) {
	user := userdb.User{Username: username}
	err := u.userRepository.CreateUserinDB(&user)
	return user, err
}
func (u UserServiceImpl) EditUser(user u.User) error {
	err := u.userRepository.EditUserinDB(user)
	if err != nil {
		return err
	}
	return nil
}

func (u UserServiceImpl) DeleteUser(idstr string) error {
	id, err := strconv.Atoi(idstr)
	if err != nil {
		
		return err
	}
	err = u.userRepository.DeleteUserinDB(id)
	if err != nil {
		return err
	}
	return nil
}

func (u UserServiceImpl) GetAllUsers() []u.User {
	users, err := u.userRepository.GetAllUsersfromDB()
	if err != nil {
		fmt.Println("Get users error:", err)
	}
	return users
}
