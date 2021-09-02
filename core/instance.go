package core

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/kelseyhightower/envconfig"
	"github.com/mitchellh/go-homedir"
	"modify-aliyun-firewall/logger"
)

type Instance struct {
	RegionId string     `toml:"region_id"`
	Firewall []FireWall `toml:"firewall"`
}

func NewInstance() *Instance {
	auth := Instance{}
	path, _ := homedir.Expand("./config.toml")

	if _, err := toml.DecodeFile(path, &auth); err != nil {
		logger.MyLogger.ErrorLog.Panic(err)
	}

	err := envconfig.Process("myapp", &auth)
	if err != nil {
		logger.MyLogger.ErrorLog.Panic(err)
	}
	return &auth
}

type FireWall struct {
	Id         string `toml:"id"`
	Name       string `toml:"name"`
	GroupName  string `toml:"group_name"`
	ModifyMode string `toml:"modify_mode"`
	Type       string `toml:"type"`
}

func (result FireWall) String() string {
	return fmt.Sprintf("instance id: %-20v\t name: %-20v\t group: %-20v\t ", result.Id, result.Name, result.GroupName)
}
