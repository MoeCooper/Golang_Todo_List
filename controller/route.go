package controller

import (
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	// (2) we register multiplexer (mux) with our endpoints "ping" in this case, calls ping()
	mux.HandleFunc("/ping", ping())
	return mux
}
