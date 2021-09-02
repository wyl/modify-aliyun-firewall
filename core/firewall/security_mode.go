package firewall

type ModifyMode string

const (
	Cover  ModifyMode = "Cover"  // 覆盖原白名单。
	Append            = "Append" // 追加白名单
	Delete            = "Delete" // 删除该白名单。
)
