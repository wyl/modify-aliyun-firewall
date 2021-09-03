package src

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dds"
	r_kvstore "github.com/aliyun/alibaba-cloud-sdk-go/services/r-kvstore"
	"log"
	"modify-aliyun-firewall/src/core"
	"modify-aliyun-firewall/src/logger"
)

type IndexClient struct {
	AliyunSecret *core.AliSecret
	Instance     *core.Instance
	Ip           string
}

func NewIndexClient(accessKeyId, accessKeySecret string) *IndexClient {

	ic := IndexClient{
		AliyunSecret: core.NewAliSecret(accessKeyId, accessKeySecret),
		Instance:     core.NewInstance(),
		Ip:           core.GetIP(),
	}

	return &ic
}

func (ic *IndexClient) Apply() {
	for _, item := range ic.Instance.Firewall {
		switch item.Type {
		case "Dds":
			if client, err := dds.NewClientWithAccessKey(ic.Instance.RegionId, ic.AliyunSecret.KeyId, ic.AliyunSecret.KeySecret); err != nil {
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
			if client, err := r_kvstore.NewClientWithAccessKey(ic.Instance.RegionId, ic.AliyunSecret.KeyId, ic.AliyunSecret.KeySecret); err != nil {
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

func (ic *IndexClient) Show() {

	logger.MyLogger.InfoLog.Println("获取到的信息如下：")
	b, _ := json.MarshalIndent(ic, "", "\t")
	fmt.Println(string(b))
}

//
//func main() {
//
//	ic := NewIndexClient()
//
//	//b, _ := json.MarshalIndent(ic.Instance, "", "\t")
//	//fmt.Println(string(b))
//	if ic.AliyunSecret.KeyId == "" || ic.AliyunSecret.KeySecret == "" {
//		logger.MyLogger.ErrorLog.Panic(fmt.Sprintf("请确认你的认证信息 KeyId:%v,KeySecret:%v", ic.AliyunSecret.KeyId, ic.AliyunSecret.KeySecret))
//	}
//
//	if ic.Show() {
//		ic.Apply()
//	}
//
//}
