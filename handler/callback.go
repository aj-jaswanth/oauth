package handler

import (
	"github.com/aj-jaswanth/oauth/oauth"
	"github.com/aj-jaswanth/oauth/store"
	"log"
	"net/http"
	"strings"
)

func Callback(w http.ResponseWriter, req *http.Request) {
	log.Println("call back is invoked")
	code := strings.Split(req.URL.RawQuery, "=")[1]
	store.SetAccessCode(oauth.AccessToken(code))
}
