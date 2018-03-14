package oauth

import (
	"bytes"
	"encoding/json"
	"github.com/aj-jaswanth/oauth/util"
	"log"
	"net/http"
)

func AccessToken(code string) string {
	client := http.DefaultClient
	request := accessTokenRequest{
		ClientId:     ClientId,
		ClientSecret: ClientSecret,
		Code:         code}
	reqInBytes, err := json.Marshal(request)
	util.HandleErr(err)
	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewReader(reqInBytes))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	util.HandleErr(err)
	resp, err := client.Do(req)
	log.Println("Requested access_token")
	tokenResp := toTokenResp(util.ResponseBodyInBytes(resp.Body))
	return tokenResp.AccessToken
}

func toTokenResp(body []byte) accessTokenResponse {
	var tokenResp accessTokenResponse
	log.Println(string(body))
	err := json.Unmarshal(body, &tokenResp)
	util.HandleErr(err)
	return tokenResp
}
