package usecases

import (
	"context"
	"ddd-template/domains/user"
	"fmt"
)

type UserUseCase struct {
	Repository user.Repository
}

func NewUserUseCase(repository user.Repository) *UserUseCase {
	return &UserUseCase{
		Repository: repository,
	}
}

func (u *UserUseCase) GetUserById(ctx context.Context, userId string) (*user.User, error) {
	usr, err := u.Repository.FindById(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("UserUseCase.GetUserById: %w", err)
	}
	if usr.Role == "admin" {
		usr.Name = fmt.Sprintf("%s:%s", usr.Role, usr.Name)
	}
	return usr, nil
}
