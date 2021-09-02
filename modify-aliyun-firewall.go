package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dds"
	r_kvstore "github.com/aliyun/alibaba-cloud-sdk-go/services/r-kvstore"
	"log"
	"modify-aliyun-firewall/core"
	"modify-aliyun-firewall/logger"
)

type IndexClient struct {
	AliyunSecret *core.AliSecret
	Instance     *core.Instance
	Ip           string
}

func NewIndexClient() *IndexClient {
	ic := IndexClient{
		AliyunSecret: core.NewAliSecret(),
		Instance:     core.NewInstance(),
		Ip:           core.GetIP(),
	}

	return &ic
}

func (ic *IndexClient) Apply() {
	for _, item := range ic.Instance.Firewall {
		switch item.Type {
		case "Dds":
			if client, err := dds.NewClientWithAccessKey(ic.Instance.RegionId, ic.AliyunSecret.AccessKeyID, ic.AliyunSecret.AccessKeySecret); err != nil {
				logger.MyLogger.ErrorLog.Panic(err)
			} else {
				params := dds.CreateModifySecurityIpsRequest()
				params.ModifyMode = string(item.ModifyMode)
				params.SecurityIps = ic.Ip
				params.SecurityIpGroupName = item.GroupName
				params.DBInstanceId = item.Id

				res, err := client.ModifySecurityIps(params)
				if err != nil {
					fmt.Println(item)
					fmt.Println(res)
					log.Println(err)
				}
				logger.MyLogger.InfoLog.Println(item, res.IsSuccess())
			}

		case "R-kvstore":
			if client, err := r_kvstore.NewClientWithAccessKey(ic.Instance.RegionId, ic.AliyunSecret.AccessKeyID, ic.AliyunSecret.AccessKeySecret); err != nil {
				logger.MyLogger.ErrorLog.Panic(err)
			} else {
				params := r_kvstore.CreateModifySecurityIpsRequest()
				params.ModifyMode = string(item.ModifyMode)
				params.SecurityIps = ic.Ip
				params.SecurityIpGroupName = item.GroupName
				params.InstanceId = item.Id
				res, err := client.ModifySecurityIps(params)
				if err != nil {
					log.Println(err)
				}
				logger.MyLogger.InfoLog.Println(item, res.IsSuccess())
			}
		}
	}
}

func (ic *IndexClient) Confirm() bool {
	fmt.Println(fmt.Sprintf("IP:%v", ic.Ip))
	fmt.Println(fmt.Sprintf("KeyID:%v/KeySecret:%v", ic.AliyunSecret.AccessKeyID, ic.AliyunSecret.AccessKeySecret))
	fmt.Println("Instance List:")
	for _, v := range ic.Instance.Firewall {
		logger.MyLogger.InfoLog.Println(v)
	}
	return logger.YesNo("")
}

func main() {

	ic := NewIndexClient()

	//b, _ := json.MarshalIndent(ic.Instance, "", "\t")
	//fmt.Println(string(b))
	if ic.AliyunSecret.AccessKeyID == "" || ic.AliyunSecret.AccessKeySecret == "" {
		logger.MyLogger.ErrorLog.Panic(fmt.Sprintf("请确认你的认证信息 AccessKeyID:%v,AccessKeySecret:%v", ic.AliyunSecret.AccessKeyID, ic.AliyunSecret.AccessKeySecret))
	}

	if ic.Confirm() {
		ic.Apply()
	}

}
