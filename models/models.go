package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	PhoneNumber string `json:"phonenumber"`
	Email       string `json:"email"`
	Sex         string `json:"sex"`
	Age         int    `json:"age"`
}

func (u *User) FillData(name, surname, phonenumber, email, sex string, age int) {
	u.Name = name
	u.Surname = surname
	u.PhoneNumber = phonenumber
	u.Email = email
	u.Sex = sex
	u.Age = age
}

func (u *User) AddToDatabase(db *gorm.DB) error {
	return db.Create(u).Error
}

type Authorization struct {
	gorm.Model
	Login    string `json:"login"`
	Password string `json:"password"`
}
