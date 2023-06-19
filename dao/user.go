package dao

import (
	"test_di/db"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type UserDao struct {
	db *db.DB
}

func (d *UserDao) GetUser() *User {
	return &User{
		Id:   123,
		Name: "name",
		Age:  1100,
	}
}

func NewUserDao(db *db.DB) *UserDao {
	return &UserDao{db: db}
}
