package config

import (
	"option_get/constants"
	"os"
)

var IsDebug = true

var ApproveKey string

func init() {
	ApproveKey = os.Getenv("APPROVE_KEY")
}

func GetAuthorizationHeader() string {
	if token := os.Getenv(constants.TenantAccessTokenKey); token != "" {
		return token
	}
	if token := os.Getenv(constants.UserAccessTokenKey); token != "" {
		return token
	}
	return ""
}
