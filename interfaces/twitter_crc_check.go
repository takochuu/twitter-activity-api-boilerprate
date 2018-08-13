package interfaces

import (
	"net/http"

	"github.com/takochuu/twitter-activity-api-boilerprate/usecase"
)

type TwitterCRCCheckHandler struct {
	TwitterCRCCheckUseCase usecase.TwitterCRCCheckInteractor
}

func NewTwitterCRCCheckHandler() *TwitterCRCCheckHandler {
	return &TwitterCRCCheckHandler{
		usecase.NewTwitterCRCCheckUseCase(),
	}
}

func (h TwitterCRCCheckHandler) Check(res http.ResponseWriter, req *http.Request) {
	if token := req.FormValue("crc_token"); len(token) > 0 {
		h.TwitterCRCCheckUseCase.Check(token)
	}
}
