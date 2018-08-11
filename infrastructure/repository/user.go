package repository

import "github.com/takochuu/go-cleanarchitecture/domain/entity"

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Create() (*entity.User, error) {
	return &entity.User{}, nil
}
