package controller

import (
	viewer "GOLANG_WEBSITE/views"
	"encoding/json"
	"net/http"
)

// (3) handler function called ping to return data
func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := viewer.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			json.NewEncoder(w).Encode(data)
		}
	}
}
