package response

import "net/http"

type Response interface {
	JSON(res http.ResponseWriter)
}
