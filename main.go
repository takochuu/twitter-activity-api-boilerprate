package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/takochuu/go-cleanarchitecture/infrastructure/repository"
	"github.com/takochuu/go-cleanarchitecture/interfaces"
	"github.com/takochuu/go-cleanarchitecture/usecase"
)

func main() {

	// TODO Router化
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		h := interfaces.UserHandler{}
		u := repository.NewUserRepository()
		h.UserUseCase = usecase.NewUserUseCase(u)
		h.Create(res, req)
	})

	http.ListenAndServe(":16000", nil)
}
