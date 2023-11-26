package userdb

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       int    `gorm:"column:id; primary_key; not null" json:"id"`
	Username string `gorm:"column:username" json:"username"`
}
type UserDB interface {
	CreateUserinDB(user *User) error
	EditUserinDB(User) error
	DeleteUserinDB(int) error
	GetAllUsersfromDB() ([]User, error)
}
type UserDBImpl struct {
	db *gorm.DB
}

func UserRepositoryInit() *UserDBImpl {
	dsn := "admin:admin@tcp(127.0.0.1:3306)/admin"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dsn == "" {
		fmt.Println("DB_DSN environment variable is not set.")

	}

	db.AutoMigrate(&User{})

	if err != nil {
		fmt.Println("db:", err)
	}

	db.AutoMigrate(&User{})
	return &UserDBImpl{
		db: db,
	}
}

func (u *UserDBImpl) CreateUserinDB(user *User) error {
	err := u.db.Save(user).Error
	return err
}

func (u *UserDBImpl) EditUserinDB(user User) error {
	
	err := u.db.Save(&user).Error

	return err
}

func (u *UserDBImpl) DeleteUserinDB(id int) error {
	user := &User{ID: id}
	err := u.db.Where("id = ?", id).Delete(&user).Error
	if err != nil {

		return err
	}
	return nil

}

func (u UserDBImpl) GetAllUsersfromDB() ([]User, error) {
	var users []User
	var err = u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
