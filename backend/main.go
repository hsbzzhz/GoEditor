package main

import (
	"editorDemo/src/handler"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/show", handler.DemoHandler)
	mux.HandleFunc("/add", handler.UploadHandler)
	// 启动静态文件服务
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/static/"))))
	mux.HandleFunc("/", handler.IndexHandler)

	http.ListenAndServe(":8080", mux)
}
