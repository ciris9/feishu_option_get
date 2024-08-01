package lark_api

import (
	"option_get/config"
	"testing"
)

func TestToken(t *testing.T) {
	token, err := config.GetAuthToken()
	if err != nil {
		t.Error(err)
	}
	println(token)
}
