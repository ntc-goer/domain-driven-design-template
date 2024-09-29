package user

import "context"

type Repository interface {
	FindById(ctx context.Context, id string) (*User, error)
}
