package setting

import (
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.AddConfigPath("configs/")
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	//vp.GetDuration("60s")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}
