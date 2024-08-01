package lark_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"option_get/config"
	"option_get/constants"
	"option_get/lark_api/api"
	"option_get/model"
)

func GetBiTableRecords(appToken, tableID string) (*model.BiTableRecordsResponse, error) {
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf(api.BiTableRecordsSearchAPI, appToken, tableID), bytes.NewBuffer([]byte(`{}`)))
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
	result := model.NewBiTableRecordsResp()
	err = json.Unmarshal(responseBody, result)
	//log.Debugf("get bitable records %#v", result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetBiTableSingleRecords(appToken, tableID, recordID string) (*model.BiTableRecordsResponse, error) {
	request, err := http.NewRequest(http.MethodPost, api.BiTableSingleRecordsSearchAPI, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", "Bearer "+config.GetAuthorizationHeader())
	request.Header.Set("Content-Type", "application/json")
	request.SetPathValue(constants.AppTokenKey, appToken)
	request.SetPathValue(constants.TableIDKey, tableID)
	request.SetPathValue(constants.RecordIDKey, recordID)
	client := http.Client{
		Timeout: 20,
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	result := model.NewBiTableRecordsResp()
	err = json.Unmarshal(responseBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetBiTableBatchRecords(appToken, tableID string, recordIDs []string) (*model.BiTableRecordsResponse, error) {
	results := model.NewBiTableRecordsResp()
	for _, recordID := range recordIDs {
		request, err := http.NewRequest(http.MethodPost, api.BiTableSingleRecordsSearchAPI, nil)
		if err != nil {
			return nil, err
		}
		request.Header.Set("Authorization", "Bearer "+config.GetAuthorizationHeader())
		request.Header.Set("Content-Type", "application/json")
		request.SetPathValue(constants.AppTokenKey, appToken)
		request.SetPathValue(constants.TableIDKey, tableID)
		request.SetPathValue(constants.RecordIDKey, recordID)
		client := http.Client{
			Timeout: 20,
		}
		response, err := client.Do(request)
		if err != nil {
			return nil, err
		}
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		result := model.NewBiTableRecordsResp()
		err = json.Unmarshal(responseBody, result)
		if err != nil {
			return nil, err
		}
		results.Data.Items = append(results.Data.Items, result.Data.Items...)
	}

	return results, nil
}
