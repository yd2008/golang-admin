package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type WechatSession struct {
	SessionKey string `json:"session_key"`
	Openid     string `json:"openid"`
}

func GetWechatAccessSession(accessToken, appID, secret, code string) (WechatSession, error) {

	url := fmt.Sprintf("%s?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", accessToken, appID, secret, code)

	response, _ := http.Get(url)
	if response.Body != nil {
		defer response.Body.Close()
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var wechatSession WechatSession
	json.Unmarshal(body, &wechatSession)

	return wechatSession, nil
}
