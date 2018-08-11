package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/takochuu/twitter-activity-api-boilerprate/interfaces"
	"github.com/takochuu/twitter-activity-api-boilerprate/usecase"
)

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		h := interfaces.TwitterCRCCheckHandler{}
		h.TwitterCRCCheckUseCase = usecase.NewTwitterCRCCheckUseCase()
		h.Check(res, req)
	})

	http.ListenAndServe(":16000", nil)
}
