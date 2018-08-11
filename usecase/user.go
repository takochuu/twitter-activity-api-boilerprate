package usecase

import "github.com/takochuu/go-cleanarchitecture/domain/entity"

type UserUseCase struct {
	repo UserInterface
}

type UserInterface interface {
	Create() (*entity.User, error)
}

func NewUserUseCase(r UserInterface) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

func (u UserUseCase) Create() (*entity.User, error) {
	user, err := u.repo.Create()
	if err != nil {
		return nil, err
	}
	return user, nil
}
