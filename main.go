package main

import (
	"github.com/aj-jaswanth/oauth/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/oauth/callback", handler.Callback)
	http.HandleFunc("/github/profile", handler.Profile)

	http.ListenAndServe("localhost:8080", nil)
}
