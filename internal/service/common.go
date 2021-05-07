package service

import (
	"golang-admin/global"
	"golang-admin/internal/model"
	"golang-admin/pkg/app"
	"golang-admin/pkg/util"
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
	session, err := app.GetWechatAccessSession(accessToken, appID, secret, code)
	if err != nil {
		return nil, err
	}

	userInfo, err := util.WechatUserInfo(param.Iv, param.EncryptedData, appID, session.SessionKey)
	if err != nil {
		return nil, err
	}

	return svc.dao.GetOrCreateUser(userInfo)
}

