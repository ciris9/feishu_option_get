package service

import (
	"encoding/json"
	"io"
	"net/http"
	"option_get/config"
	"option_get/constants"
	"option_get/model"
)

func GetBiTableRecords(appToken, tableID string) (*model.BiTableResponse, error) {
	request, err := http.NewRequest(http.MethodPost, constants.BiTableRecordsSearchAPI, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", "Bearer "+config.GetAuthorizationHeader())
	request.Header.Set("Content-Type", "application/json")
	request.SetPathValue(constants.AppTokenKey, appToken)
	request.SetPathValue(constants.TableIDKey, tableID)
	client := http.Client{
		Timeout: 20,
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	responseBody, err := io.ReadAll(response.Body)
	result := model.NewBiTableResp()
	err = json.Unmarshal(responseBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetBiTableSingleRecords(appToken, tableID, recordID string) (*model.BiTableResponse, error) {
	request, err := http.NewRequest(http.MethodPost, constants.BiTableSingleRecordsSearchAPI, nil)
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
	result := model.NewBiTableResp()
	err = json.Unmarshal(responseBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
	{
	  "code": 0,
	  "data": {
	    "absent_record_ids": [],
	    "forbidden_record_ids": [],
	    "records": [
	      {
	        "fields": {
	          "Multiple Options": [
	            "FASD",
	            "ZXCAS"
	          ],
	          "Number": 123,
	          "人员": [
	            {
	              "avatar_url": "https://s1-imfile.feishucdn.com/static-resource/v1/v3_00d4_07f9e58b-b7b7-4008-8972-cc4da428303g~?image_size=72x72&cut_type=default-face&quality=&format=jpeg&sticker_format=.webp",
	              "email": "",
	              "en_name": "余啟旺",
	              "id": "ou_73d32f8b615a4f0ce19ff5a846d8e27d",
	              "name": "余啟旺"
	            }
	          ],
	          "单选": "AS",
	          "文本": [
	            {
	              "text": "asdfafrs",
	              "type": "text"
	            }
	          ],
	          "日期": 1721923200000
	        },
	        "record_id": "rec1UMKrxD"
	      }
	    ]
	  },
	  "msg": "success"
	}
*/
func GetBiTableBatchRecords(appToken, tableID string, recordIDs []string) (*model.BiTableResponse, error) {
	results := model.NewBiTableResp()
	for _, recordID := range recordIDs {
		request, err := http.NewRequest(http.MethodPost, constants.BiTableSingleRecordsSearchAPI, nil)
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
		result := model.NewBiTableResp()
		err = json.Unmarshal(responseBody, result)
		if err != nil {
			return nil, err
		}
		results.Data.Items = append(results.Data.Items, result.Data.Items...)
	}

	return results, nil
}
