package service

import (
	"test_di/dao"
)

type UserService struct {
	dao *dao.UserDao
}

func (svc *UserService) GetUser() *dao.User {
	return svc.dao.GetUser()
}

func NewUserService(dao *dao.UserDao) *UserService {
	return &UserService{dao: dao}
}
