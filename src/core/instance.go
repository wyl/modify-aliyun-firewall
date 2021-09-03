package core

import (
	"fmt"
	"github.com/spf13/viper"
	firewall2 "modify-aliyun-firewall/src/core/firewall"
)

type Instance struct {
	RegionId string
	Firewall []FireWall
}

func NewInstance() *Instance {
	auth := Instance{}
	viper.Unmarshal(&auth)
	return &auth
}

type FireWall struct {
	Id         string
	Name       string
	GroupName  string
	ModifyMode firewall2.ModifyMode
	Type       string
}

func (result FireWall) String() string {
	return fmt.Sprintf("InstanceId: %-20v\t %-20v\t %-20v\t %-20v\t ", result.Id, result.Name, result.GroupName, result.ModifyMode)
}
