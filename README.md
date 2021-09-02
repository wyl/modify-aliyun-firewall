

# Modify Aliyun Firewall

### Aliyun secret
```shell
export ALIYUN_ACCESS_KEY_ID=YOURID
export ALIYUN_ACCESS_KEY_SECRET=YOUR_SECRET
```

### 更改的 Instance IDs
```toml

# This is a TOML document. Boom.
region_id = "cn-beijing"

[[firewall]]
id = "dds-instance-id"
name = "dds-description-name"
group_name = "shenzhen"
modify_mode = "Cover"
type = "Dds"

[[firewall]]
id = "r-instance-id"
name = "r-description-name"
group_name = "shenzhen"
modify_mode = "Cover"
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

### 执行
```shell script

./modify-aliyun-firewall 
```
