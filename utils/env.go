package utils

import (
	"github.com/spf13/viper"
)

var EnvConfig *Env

type Env struct {
	GoNode string `mapstructure:"GO_NODE"`
}

func InitEnv() error {
	EnvConfig = &Env{}

	viper.SetDefault("GO_NODE", "0")
	viper.AutomaticEnv() // 自動讀取環境變數

	// 將配置綁定到結構體
	if err := viper.Unmarshal(EnvConfig); err != nil {
		return err
	}

	return nil
}
