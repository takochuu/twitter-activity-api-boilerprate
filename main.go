package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/takochuu/twitter-activity-api-boilerprate/interfaces"
)

func main() {

	http.HandleFunc("/crc_token", func(res http.ResponseWriter, req *http.Request) {
		h := interfaces.NewTwitterCRCCheckHandler()
		h.GenerateCRCToken(res, req)
	})

	http.ListenAndServe(":16000", nil)
}
