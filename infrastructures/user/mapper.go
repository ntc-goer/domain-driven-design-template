package user

import (
	"ddd-template/domains/user"
	"ddd-template/infrastructures/gorm/userdao"
)

func mapDaoToDomain(userDao *userdao.User) *user.User {
	return &user.User{
		Id:        userDao.Id,
		Name:      userDao.Name,
		Role:      userDao.Role,
		Address:   userDao.Address,
		CreatedAt: userDao.CreatedAt,
		UpdatedAt: userDao.UpdatedAt,
	}
}
