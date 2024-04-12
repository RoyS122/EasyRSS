package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)

	switch status {
	case 404:
		t, err := template.ParseFiles("rsc/html/404.html")
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, nil)

	case 400:
		t, err := template.ParseFiles("rsc/html/400.html")
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, nil)

	default:
		t, err := template.ParseFiles("rsc/html/500.html")
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, nil)
	}

}