package main

import (
	"editorDemo/src/handler"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/golang/run-code", handler.UploadHandler)
	// build server for static files for frontend
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/static/"))))
	mux.HandleFunc("/", handler.IndexHandler)

	http.ListenAndServe(":8080", mux)
}
