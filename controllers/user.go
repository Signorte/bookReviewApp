package controllers

import (
	"github.com/Signorte/bookReviewApp/models"
)

// User is
type User struct {
}

// NewUser ...
func NewUser() User {
	return User{}
}

// Get ...
func (c User) Get(n int) interface{} {
	repo := models.NewUserRepository()
	user := repo.GetByID(n)
	//user_json := params.User{
	//ID:       user.ID,
	//Username: user.Username,
	//Bookname: user.Bookname,
	//Review:   user.Review,
	//}
	//return user
	return user
}
