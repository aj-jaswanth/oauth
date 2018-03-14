package handler

import (
	"bytes"
	"encoding/json"
	"github.com/aj-jaswanth/oauth/store"
	"github.com/aj-jaswanth/oauth/util"
	"log"
	"net/http"
	"strings"
)

type user struct {
	URL string `json:"blog"`
}

func Profile(w http.ResponseWriter, req *http.Request) {
	url := strings.Split(req.URL.RawQuery, "=")[1]
	log.Printf("url to set to: %s", url)
	client := http.DefaultClient
	updateReq, err := http.NewRequest("PATCH", "https://api.github.com/user", bytes.NewReader(userInBytes(url)))
	util.HandleErr(err)
	updateReq.Header.Set("Accept", "application/vnd.github.v3+json")
	updateReq.Header.Set("Content-Type", "application/json")
	updateReq.Header.Set("Authorization", "token "+store.GetAccessCode())
	res, err := client.Do(updateReq)
	util.HandleErr(err)
	log.Printf("got response: %d", res.StatusCode)
}

func userInBytes(url string) []byte {
	user := user{URL: url}
	userInBytes, err := json.Marshal(user)
	util.HandleErr(err)
	return userInBytes
}
