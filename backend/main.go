package main

import (
	"editorDemo/src/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.DemoHandler)
	http.HandleFunc("/add", handler.UploadHandler)
	http.ListenAndServe(":8080", nil)
}
