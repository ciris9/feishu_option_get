package lark_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"option_get/config"
	"option_get/lark_api/api"
	"option_get/model"
)

func GetBiTableFields(appToken, tableID string) (*model.BiTableFieldResponse, error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf(api.BiTableFieldsListAPI, appToken, tableID), bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", "Bearer "+config.Token)
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	result := model.NewBITableFieldResponse()
	err = json.Unmarshal(responseBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
