package dao

import (
	"golang-admin/internal/model"
	"golang-admin/pkg/util"
)

func (d *Dao) GetOrCreateUser(wechatUser *util.WechatUser) (*model.User, error) {
	var user = model.User{
		IsWechatLogin: 1,
		WechatId: wechatUser.OpenID,
		Username: wechatUser.NickName,
		Gender: uint8(wechatUser.Gender),
		Avatar: wechatUser.AvatarURL,
	}
	return user.WechatLogin(d.engine)
}
