package config

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"option_get/constants"
	"option_get/lark_api/api"
	"option_get/model"
	"os"
	"time"
)

var IsDebug = true

var Token string

var ApproveKey string

func init() {
	ApproveKey = os.Getenv("APPROVE_KEY")
	Token, _ = GetAuthToken()
	go func() {
		DynamicGetToken()
	}()
}

func GetAuthToken() (token string, err error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	tenantAccessTokenRequest := model.NewTenantAccessTokenRequest()
	tenantAccessTokenRequest.AppSecret = GetAppSecret()
	tenantAccessTokenRequest.AppID = GetAppID()
	body, err := json.Marshal(tenantAccessTokenRequest)
	if err != nil {
		return "", err
	}
	println(string(body))
	request, err := http.NewRequest(http.MethodPost, api.AccessTokenGetAPI, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	result := model.NewTenantAccessTokenResponse()
	tokenResponse, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	if err = json.Unmarshal(tokenResponse, result); err != nil {
		return "", err
	}
	log.Println(result.TenantAccessToken)
	return result.TenantAccessToken, nil
}

func GetAuthorizationHeader() string {
	if IsDebug {
		return "t-g1047teeOTXGP2YT6CZ35JYEAFMDCMR4SDIP4PBZ"
	}
	if token := os.Getenv(constants.TenantAccessTokenKey); token != "" {
		return token
	}
	if token := os.Getenv(constants.UserAccessTokenKey); token != "" {
		return token
	}
	return ""
}

func DynamicGetToken() {
	var err error
	tick := time.NewTicker(time.Minute * 30)
	for range tick.C {
		Token, err = GetAuthToken()
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
func GetAppID() string {
	if IsDebug {
		return "cli_a62b0a38a77ad00e"
	}
	return os.Getenv(constants.AppIDKey)
}

func GetAppSecret() string {
	if IsDebug {
		return "G3TPCEOJxkqWPP5rgbcHWdGvX8SqyvcD"
	}
	return os.Getenv(constants.AppSecretKey)
}
