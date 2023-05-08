package main

import (
	"net/http"

	"github.com/DB-Vincent/k8s-demo-app/html"
)

func main() {
	http.HandleFunc("/env", env)
	http.ListenAndServe(":8080", nil)
}

func env(w http.ResponseWriter, r *http.Request) {
	p := html.EnvParams{
		Title:   "Some super cool title",
		Message: "Hello there!",
	}
	html.Env(w, p)
}