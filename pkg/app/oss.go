package app

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"golang-admin/global"
)

type OssCredentials struct {
	AccessKeySecret string `json:"access_key_secret"`
	Expiration      string `json:"expiration"`
	AccessKeyId     string `json:"access_key_id"`
	SecurityToken   string `json:"security_token"`
}

func CreateCredentials() (*OssCredentials, error) {
	ossSetting := global.ALiOSSSetting
	//构建一个阿里云客户端, 用于发起请求。
	//构建阿里云客户端时，需要设置AccessKey ID和AccessKey Secret。
	client, err := sts.NewClientWithAccessKey(ossSetting.RegionId, ossSetting.AccessKeyId, ossSetting.AccessKeySecret)

	//构建请求对象。
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = ossSetting.Scheme

	//设置参数。关于参数含义和设置方法，请参见API参考。
	request.RoleArn = ossSetting.RoleArn
	request.RoleSessionName = ossSetting.RoleSessionName

	//发起请求，并得到响应。
	response, err := client.AssumeRole(request)
	if err != nil {
		return nil, err
	}
	var ossCredentials = &OssCredentials{
		AccessKeySecret: response.Credentials.AccessKeySecret,
		Expiration:      response.Credentials.Expiration,
		AccessKeyId:     response.Credentials.AccessKeyId,
		SecurityToken:   response.Credentials.SecurityToken,
	}

	return ossCredentials, nil
}
