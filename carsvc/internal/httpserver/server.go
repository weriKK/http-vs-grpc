package httpserver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type httpserver struct {
}

func NewHTTPServer(port int) *http.Server {

	h := httpserver{}

	mux := mux.NewRouter()
	mux.HandleFunc("/car", h.carDataHandler).Methods("POST")

	return &http.Server{
		Addr:    fmt.Sprintf("127.0.0.1:%d", port),
		Handler: mux,
	}
}
