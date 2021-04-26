package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
)

type WechatUser struct {
	OpenID    string `json:"openId"`
	UnionID   string `json:"unionId"`
	NickName  string `json:"nickName"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarURL string `json:"avatarUrl"`
	Language  string `json:"language"`
	Watermark struct {
		Timestamp int64  `json:"timestamp"`
		AppID     string `json:"appid"`
	} `json:"watermark"`
}

func WechatUserInfo(iv, encryptedData, appID, sessionKey string) (*WechatUser, error) {
	// base64 解密参数
	_iv, err1 := base64.StdEncoding.DecodeString(iv)
	_encryptedData, err2 := base64.StdEncoding.DecodeString(encryptedData)
	_sessionKey, err3 := base64.StdEncoding.DecodeString(sessionKey)
	if err1 != nil || err2 != nil || err3 != nil {
		return nil, errors.New("解密失败")
	}
	var userData WechatUser
	// aes-128-cbc 解密
	cipherBlock, err := aes.NewCipher(_sessionKey)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(cipherBlock, _iv)
	mode.CryptBlocks(_encryptedData, _encryptedData)
	// 填充
	decrypted := unPad(_encryptedData)
	err = json.Unmarshal(decrypted, &userData)
	if err != nil {
		return nil, err
	}

	if userData.Watermark.AppID != appID {
		return nil, errors.New("appid 不一致！")
	}
	return &userData, nil
}

// 以PKCS#7填充方式的填充方法
func unPad(s []byte) []byte {
	return s[:(len(s) - int(s[len(s)-1]))]
}
