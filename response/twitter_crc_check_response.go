package response

type TwitterCRCCheckResponse struct {
	CRCToken string `json:"crc_token"`
}

func NewTwitterCRCCheckResponse() *Response {
	return &TwitterCRCCheckResponse{}
}

func (r *TwitterCRCCheckResponse) JSON(data interface{}) {
}

func (r *TwitterCRCCheckResponse) GRPC(data interface{}) {
}
