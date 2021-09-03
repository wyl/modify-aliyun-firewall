

# Modify Aliyun Firewall

### file 
配置文件参考，_config.toml
### modify-aliyun-firewall
目录信息


```shell script
wangyalong@wangyalongde-MacBook-Pro-for-Job modify-aliyun-firewall_0.0.2_Linux_amd64$ tree -l
.
├── config.toml
└── modify-aliyun-firewall

0 directories, 2 files

```

配置文件config.toml  参考_config.toml
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
默认配置文件 `.config.toml` ，自定义配置文件 --config <your-config>。

### example 
```shell script
./modify-aliyun-firewall show --config <your-config>
./modify-aliyun-firewall apply --config <your-config>
```



### Aliyun secret
```shell
export ALIYUN_ACCESS_KEY_ID=YOURID
export ALIYUN_ACCESS_KEY_SECRET=YOUR_SECRET
```
