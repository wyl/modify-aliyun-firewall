package core

import (
	"github.com/kelseyhightower/envconfig"
	"modify-aliyun-firewall/logger"
)

type AliSecret struct {
	AccessKeyID     string `envconfig:"ALIYUN_ACCESS_KEY_ID"`
	AccessKeySecret string `envconfig:"ALIYUN_ACCESS_KEY_SECRET"`
}

func NewAliSecret() *AliSecret {
	auth := AliSecret{}
	err := envconfig.Process("myapp", &auth)
	if err != nil {
		logger.MyLogger.ErrorLog.Panic(err)
	}

	return &auth
}
