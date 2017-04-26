package deke

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

var server *httptest.Server

func RespondWith(s string) {
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, s)
	}))
}

func Close() {
	server.Close()
}

func Url() string {
	return server.URL
}
