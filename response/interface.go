package response

import "net/http"

type ResponseInterface interface {
	JSON(res http.ResponseWriter)
}
