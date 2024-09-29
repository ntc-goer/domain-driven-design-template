package user

import (
	"context"
	"ddd-template/domains/user"
	"ddd-template/infrastructures/gorm/userdao"
)

type Repository struct {
	UserDao *userdao.Dao
}

func NewRepository(userDao *userdao.Dao) *Repository {
	return &Repository{
		UserDao: userDao,
	}
}

func (r *Repository) FindById(_ context.Context, id string) (*user.User, error) {
	usr, err := r.UserDao.FindById(id)
	if err != nil {
		return nil, err
	}
	return mapDaoToDomain(usr), nil
}
