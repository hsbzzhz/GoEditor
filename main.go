package main

import (
	"editorDemo/src/handle"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle.DemoHandler)
	http.ListenAndServe(":8080", nil)
}
