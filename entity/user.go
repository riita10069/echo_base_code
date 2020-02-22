package entity

import ("github.com/jinzhu/gorm"
	"github.com/riita10069/echo_base_code/db"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
}

func CreateUser(user *User) {
	db.New().Create(user)
}

func FindUser(u *User) User {
	var user User
	db.New().Where(u).First(&user)
	return user
}