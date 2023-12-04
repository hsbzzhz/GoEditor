package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/index.html")
	if err != nil {
		fmt.Println(err)
	}

	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}

	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}
