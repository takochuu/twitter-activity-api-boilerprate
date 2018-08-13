package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/takochuu/twitter-activity-api-boilerprate/interfaces"
)

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		h := interfaces.NewTwitterCRCCheckHandler()
		h.Check(res, req)
	})

	http.ListenAndServe(":16000", nil)
}
