package interfaces

import "net/http"

type TwitterCRCCheckUseCase interface {
	Check() error
}

type TwitterCRCCheckHandler struct {
	TwitterCRCCheckUseCase TwitterCRCCheckUseCase
}

func (h TwitterCRCCheckHandler) Check(res http.ResponseWriter, req *http.Request) {
	h.TwitterCRCCheckUseCase.Check()
}
