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

func UserRepositoryInit(db *gorm.DB) *UserDBImpl {
	db.AutoMigrate(&User{})
	return &UserDBImpl{
		db: db,
	}
}

func (u *UserDBImpl) CreateUserinDB(user *User) error {
	err := u.db.Save(user).Error
	return err
}

func (u *UserDBImpl) EditUserinDB(username string) error {
	user := User{
		Username: username,
	}
	err := u.db.Preload("User").First(&user).Error

	return err
}

func (u *UserDBImpl) DeleteUserinDB(username string) error {
	err := u.db.Delete(&User{}, username).Error
	if err != nil {

		return err
	}
	return nil

}

func (u UserDBImpl) GetAllUsersfromDB() ([]User, error) {
	var users []User
	var err = u.db.Preload("User").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
