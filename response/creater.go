package response

type ResponseCreater struct {
	Handler ResponseInterface
}

func NewResponseCreater(r ResponseInterface) *ResponseCreater {
	return &ResponseCreater{
		Handler: r,
	}
}
