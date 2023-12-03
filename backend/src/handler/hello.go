package handler

import (
	"fmt"
	"net/http"
)

func DemoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, x-requested-with")
	fmt.Fprintln(w, "Hello, WorldGO!")
	w.Write([]byte("dddddddd"))
}
