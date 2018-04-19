package controllers

import (
	"github.com/Signorte/bookReviewApp/models"
	"github.com/go-xorm/xorm"
	"github.com/Signorte/bookReviewApp/params"
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
	user_json = params.User(id: user.id, nickname: user.nickname)
	//return user
	return user_json
}
