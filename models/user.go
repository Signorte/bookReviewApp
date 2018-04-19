package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

// init ...
func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:1234@tcp(127.0.0.1:3306)/bookshelf")
	if err != nil {
		panic(err)
	}
}

// User is
type User struct {
	ID       int    `json:"id" xorm:"'id'"`
	Username string `json:"name" xorm:"'nickname'"`
}

// NewUser ...
func NewUser(id int, username string) User {
	return User{
		ID:       id,
		Username: username,
	}
}

// UserRepository is
type UserRepository struct {
}

// NewUserRepository ...
func NewUserRepository() UserRepository {
	return UserRepository{}
}

// GetByID ...
func (m UserRepository) GetByID(id int) *User {
	var user = User{ID: id}
	has, _ := engine.Get(&user)
	if has {
		return &user
	}
	return nil
}
