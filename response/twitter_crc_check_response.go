package response

import (
	"encoding/json"
	"net/http"
)

type TwitterCRCCheckResponse struct {
	ResponseToken string `json:"response_token"`
}

func NewTwitterCRCCheckResponse() *TwitterCRCCheckResponse {
	return &TwitterCRCCheckResponse{}
}

func (r *TwitterCRCCheckResponse) JSON(res http.ResponseWriter) {
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(r)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Write(b)
}

func (r *TwitterCRCCheckResponse) SetResponseToken(token string) {
	r.ResponseToken = token
}
