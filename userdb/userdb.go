package userdb

import "gorm.io/gorm"

type User struct {
	ID       int    `gorm:"column:id; primary_key; not null" json:"id"`
	Username string `gorm:"column:username" json:"username"`
}
type UserDB interface {
	CreateUserinDB(username string) error
	EditUserinDB(string) error
	DeleteUserinDB(string) error
	GetAllUsersfromDB() ([]User, error)
}
type UserDBImpl struct {
	db *gorm.DB
}

func (u *UserDBImpl) CreateUserinDB(user *User) {
	u.db.Save(user)
}

func (u *UserDBImpl)EditUserinDB() {

}

func (u *UserDBImpl)DeleteUserinDB() {

}

func (u UserDBImpl) GetAllUsersfromDB() ([]User, error) {
	var users []User
	var err = u.db.Preload("Role").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
