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
	ID       int    `xorm:"'id'"`
	Username string `xorm:"'nickname'"`
	Bookname string `xorm:"'bookname'"`
	Review   string `xorm:"'review'"`
}

// NewUser ...
func NewUser(id int, username string, bookname string, review string) User {
	return User{
		ID:       id,
		Username: username,
		Bookname: bookname,
		Review:   review,
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
