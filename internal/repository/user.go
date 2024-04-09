package repository

import (
	"context"
)

type UserRepository interface {
	getUserById(id int64) error
}

type userRepo struct {
	*Repository
}

func NewUserRepository(repository *Repository) UserRepository {
	return &userRepo{
		Repository: repository,
	}
}

func (ur *userRepo) getUserById(id int64) error {
	ctx := context.Background()
	ur.queries.GetUserById(ctx, id)
	return nil
}
