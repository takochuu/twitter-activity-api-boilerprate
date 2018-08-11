package interfaces

import "net/http"

type TwitterCRCCheckUseCase interface {
	Check(string) error
}

type TwitterCRCCheckHandler struct {
	TwitterCRCCheckUseCase TwitterCRCCheckUseCase
}

func (h TwitterCRCCheckHandler) Check(res http.ResponseWriter, req *http.Request) {
	if token := req.FormValue("crc_token"); len(token) > 0 {
		h.TwitterCRCCheckUseCase.Check(token)
	}
}
