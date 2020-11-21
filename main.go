package main

import (
	controllerer "GOLANG_WEBSITE/controller"
	"net/http"
)

func main() {
	// 1) entry point, routes
	mux := controllerer.Register()
	http.ListenAndServe("localhost:3009", mux)
}
