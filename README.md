

# Modify Aliyun Firewall


## help

```shell script

wangyalong@wangyalongde-MacBook-Pro-for-Job modify-aliyun-firewall_darwin_amd64$ ./modify-aliyun-firewall help

更改阿里云云产品的白名单:

目前支持：mongo/redis 服务的白名单更改。

Usage:
  modify-aliyun-firewall [command]

Available Commands:
  apply       执行更改操作
  help        Help about any command
  show        仅显示获取到的数据变量

Flags:
      --config string   config file (default is ./.config.toml)
  -h, --help            help for modify-aliyun-firewall

Use "modify-aliyun-firewall [command] --help" for more information about a command.
```


### show 
```
wangyalong@wangyalongde-MacBook-Pro-for-Job modify-aliyun-firewall_darwin_amd64$ ./modify-aliyun-firewall show
INFO 2021/09/03 15:53:00 获取到的信息如下：
{
	"AliyunSecret": {
		"KeyId": "your-id",
		"KeySecret": "our-secret"
	},
	"Instance": {
		"RegionId": "cn-beijing",
		"Firewall": [
			{
				"Id": "dds-your-id",
				"Name": "fm-ddd-dev-001",
				"GroupName": "group-name",
				"ModifyMode": "Cover",
				"Type": "Dds"
			}
		]
	},
	"Ip": "113.204.*.*"
}
```

### Aliyun secret
```shell
export ALIYUN_ACCESS_KEY_ID=YOURID
export ALIYUN_ACCESS_KEY_SECRET=YOUR_SECRET
```

### DB Instance
支持 Mongo/Redis.
```toml

# This is a TOML document. Boom.
regionId = "cn-beijing"

[[firewall]]
id = "dds-instance-id"
name = "dds-description-name"
group_name = "shenzhen"
modify_mode = "Cover"
type = "Dds"

[[firewall]]
id = "r-instance-id"
name = "r-description-name"
groupName = "shenzhen"
modifyMode = "Cover"
type = "R-kvstore"

```

### Options

- 阿里云 Secret ENV。（ALIYUN_ACCESS_KEY_ID，ALIYUN_ACCESS_KEY_SECRET）
- Instance 配置文件。 ./config.toml

目录为：
```shell script
wangyalong@wangyalongde-MacBook-Pro-for-Job modify-aliyun-firewall_0.0.2_Linux_amd64$ tree -l
.
├── config.toml
└── modify-aliyun-firewall

0 directories, 2 files

```
