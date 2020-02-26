package entity

import (
  "github.com/riita10069/echo_base_code/db"
  "time"
)

type User struct {
  ID         int64     `json:"id" gorm:"column:id;primary_key"`
	Name       string    `json:"name" gorm:"not null"`
	Password   string    `json:"password" gorm:"not null"`
  CreatedAt  time.Time `json:"created_at"`
  UpdatedAt  time.Time `json:"updated_at`
}

func CreateUser(user *User) {
	db.New().Create(user)
}

func FindUser(u *User) User {
	var user User
	db.New().Where(u).First(&user)
	return user
}
