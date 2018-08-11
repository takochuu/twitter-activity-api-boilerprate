package interfaces

import (
	"net/http"

	"github.com/takochuu/go-cleanarchitecture/domain/entity"
)

type UserUseCase interface {
	Create() (*entity.User, error)
}

type UserHandler struct {
	UserUseCase UserUseCase
}

func (h UserHandler) Create(res http.ResponseWriter, req *http.Request) {
	h.UserUseCase.Create()
}
