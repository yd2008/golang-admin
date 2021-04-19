package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	//viper.SetConfigFile("./configs/config.yaml") // 指定配置文件路径
	vp.AddConfigPath("configs/") // 查找配置文件所在的路径
	vp.SetConfigName("config")   // 配置文件名称(无扩展名)
	vp.SetConfigType("yaml")     // 如果配置文件的名称中没有扩展名，则需要配置此项
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}
