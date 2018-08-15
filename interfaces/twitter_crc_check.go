package interfaces

import (
	"net/http"

	"github.com/takochuu/twitter-activity-api-boilerprate/response"
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

func (h TwitterCRCCheckHandler) GenerateCRCToken(res http.ResponseWriter, req *http.Request) {
	if t := req.FormValue("crc_token"); len(t) > 0 {
		token := h.TwitterCRCCheckUseCase.GenerateCRCToken(t)
		r := response.NewTwitterCRCCheckResponse()
		r.SetResponseToken(token)
		r.JSON(res)
	}
}
