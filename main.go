package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/gorilla/mux"
	"github.com/takochuu/twitter-activity-api-boilerprate/interfaces"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/crc_check", func(res http.ResponseWriter, req *http.Request) {
		h := interfaces.NewTwitterCRCCheckHandler()
		h.GenerateCRCToken(res, req)
	}).Methods("GET")

	http.ListenAndServe(":8091", r)
}
