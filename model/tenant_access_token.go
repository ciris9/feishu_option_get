package model

type TenantAccessTokenRequest struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

func NewTenantAccessTokenRequest() *TenantAccessTokenRequest {
	return &TenantAccessTokenRequest{}
}

type TenantAccessTokenResponse struct {
	Code              int    `json:"app_id"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
	Expire            int64  `json:"expire"`
}

func NewTenantAccessTokenResponse() *TenantAccessTokenResponse {
	return &TenantAccessTokenResponse{}
}
