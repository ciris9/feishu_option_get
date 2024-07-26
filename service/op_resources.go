package service

import (
	"github.com/segmentio/ksuid"
	"option_get/constants"
	"option_get/model"
	"option_get/utils/datastruct/set"
	"strconv"
)

/*
	router
	OperationsResources.Post("/project_name")
	OperationsResources.Post("/resources_platform")
	OperationsResources.Post("/apply_type")
	OperationsResources.Post("/environment")
*/

/*
{
  "code": 0,
  "data": {
    "has_more": false,
    "items": [
      {
        "fields": {
          "资源平台": [
            "华为云",
            "亚马逊云"
          ],
          "项目名称": [
            {
              "text": "AIGC",
              "type": "text"
            }
          ]
        },
        "record_id": "recujCbuKbVZVE"
      },
      {
        "fields": {
          "资源平台": [
            "阿里云",
            "微软云"
          ],
          "项目名称": [
            {
              "text": "AT7",
              "type": "text"
            }
          ]
        },
        "record_id": "recujCbvZS9CdO"
      },
      {
        "fields": {
          "资源平台": [
            "华为云",
            "亚马逊云",
            "阿里云"
          ],
          "项目名称": [
            {
              "text": "AT6",
              "type": "text"
            }
          ]
        },
        "record_id": "recujCbw9yfADu"
      },
      {
        "fields": {
          "资源平台": [
            "华为云",
            "亚马逊云",
            "阿里云",
            "微软云"
          ],
          "项目名称": [
            {
              "text": "ADX",
              "type": "text"
            }
          ]
        },
        "record_id": "recujCbxJRJbXV"
      }
    ],
    "total": 4
  },
  "msg": "success"
}
*/

type OpResourcesService struct {
}

func NewOpResourcesService() *OpResourcesService {
	return &OpResourcesService{}
}

func (or *OpResourcesService) GetProjectNameOptions(request *model.ApproveRequest) (*model.ApproveResponse, error) {
	appToken := request.LinkageParams.KV[constants.AppTokenKey]
	tableToken := request.LinkageParams.KV[constants.AppTokenKey]
	records, err := GetBiTableRecords(appToken, tableToken)
	if err != nil {
		return nil, err
	}
	response := model.NewApproveResponse()
	for _, item := range records.Data.Items {
		for _, field := range item.Fields {
			fieldMap := field.(map[string]any)
			values := fieldMap[constants.ProjectNameKey].([]map[string]string)
			for i, value := range values {
				response.Data.Result.Options = append(response.Data.Result.Options, &model.ApproveOption{
					IsDefault: false,
					ID:        "option_" + strconv.Itoa(i) + "_id_" + ksuid.New().String(),
					Value:     value[constants.TextKey],
				})
			}
		}
	}
	universalApproveResponseInit(response)
	return response, nil
}

func (or *OpResourcesService) GetResourcesPlatformOptions(request *model.ApproveRequest) (*model.ApproveResponse, error) {
	appToken := request.LinkageParams.KV[constants.AppTokenKey]
	tableToken := request.LinkageParams.KV[constants.AppTokenKey]
	records, err := GetBiTableRecords(appToken, tableToken)
	if err != nil {
		return nil, err
	}
	s := set.New[string]()
	response := model.NewApproveResponse()
	for _, item := range records.Data.Items {
		for _, field := range item.Fields {
			fieldMap := field.(map[string]any)
			values := fieldMap[constants.ResourcesPlatformKey].([]string)
			for _, value := range values {
				s.Add(value)
			}
		}
	}
	var i = 0
	s.Range(func(value string) bool {
		response.Data.Result.Options = append(response.Data.Result.Options, &model.ApproveOption{
			IsDefault: false,
			ID:        "option_" + strconv.Itoa(i) + "_id_" + ksuid.New().String(),
			Value:     value,
		})
		i++
		return true
	})
	return response, nil
}

func (or *OpResourcesService) GetApplyTypeOptions(request *model.ApproveRequest) (*model.ApproveResponse, error) {
	appToken := request.LinkageParams.KV[constants.AppTokenKey]
	tableToken := request.LinkageParams.KV[constants.AppTokenKey]
	_, err := GetBiTableRecords(appToken, tableToken)
	if err != nil {
		return nil, err
	}
	response := model.NewApproveResponse()

	return response, nil
}

func (or *OpResourcesService) GetEnvOptions(request *model.ApproveRequest) (*model.ApproveResponse, error) {
	appToken := request.LinkageParams.KV[constants.AppTokenKey]
	tableToken := request.LinkageParams.KV[constants.AppTokenKey]
	_, err := GetBiTableRecords(appToken, tableToken)
	if err != nil {
		return nil, err
	}
	response := model.NewApproveResponse()

	return response, nil
}
