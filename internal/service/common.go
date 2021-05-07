package service

import (
	"encoding/json"
	"fmt"
	"golang-admin/global"
	"golang-admin/internal/model"
	"golang-admin/pkg/util"
	"io/ioutil"
	"log"
	"net/http"
)

type WehatLoginBody struct {
	Code          string `json:"code" binging:"require"`
	EncryptedData string `json:"encrypted_data" binging:"require"`
	Iv            string `json:"iv" binging:"require"`
}

func (svc *Service) WechatLogin(param *WehatLoginBody) (*model.User, error) {
	code := param.Code
	appID := global.ThirdSetting.WechatAppID
	secret := global.ThirdSetting.WechatAppSecret
	accessToken := global.ThirdSetting.WechatAccessToken
	session, err := getWechatAccessSession(accessToken, appID, secret, code)
	if err != nil {
		return nil, err
	}

	userInfo, err := util.WechatUserInfo(param.Iv, param.EncryptedData, appID, session.SessionKey)
	if err != nil {
		return nil, err
	}

	return svc.dao.GetOrCreateUser(userInfo)
}

type WechatSession struct {
	SessionKey string `json:"session_key"`
	Openid     string `json:"openid"`
}

func getWechatAccessSession(accessToken, appID, secret, code string) (WechatSession, error) {

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
