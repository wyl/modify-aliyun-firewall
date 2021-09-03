package core

type AliSecret struct {
	KeyId     string `envconfig:"ALIYUN_ACCESS_KEY_ID"`
	KeySecret string `envconfig:"ALIYUN_ACCESS_KEY_SECRET"`
}

func NewAliSecret(id, secret string) *AliSecret {
	auth := AliSecret{
		id,
		secret,
	}
	//err := envconfig.Process("myapp", &auth)
	//if err != nil {
	//	logger.MyLogger.ErrorLog.Panic(err)
	//}

	return &auth
}
