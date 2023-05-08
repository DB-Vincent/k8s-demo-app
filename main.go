package main

import (
	"net/http"

	"github.com/DB-Vincent/k8s-demo-app/web"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := web.EnvParams{
		  Title:    "Some cool title",
		  Version:  "v1.27.1"
		}

		web.Env(w, p)
	})
	http.ListenAndServe(":8080", nil)
}