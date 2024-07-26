package constants

// api
const (
	BiTableRecordsSearchAPI       = "https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/search"
	BiTableSingleRecordsSearchAPI = "https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/:record_id"
)

// auth header token key
const (
	TenantAccessTokenKey = "tenant_access_token"
	UserAccessTokenKey   = "user_access_token"
)

// params
const (
	AppTokenKey = "app_token"
	TableIDKey  = "table_id"
	RecordIDKey = "record_id"
	OptionKey   = "name"
)

const (
	ZhCNLocaleKey = "zh_cn"
	EnUSLocalKey  = "en_us"
)

const (
	TextKey = "text"
)

const (
	ProjectNameKey       = "项目名称"
	ResourcesPlatformKey = "资源平台"
)
