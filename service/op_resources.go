package service

import (
	"errors"
	"github.com/segmentio/ksuid"
	"option_get/constants"
	"option_get/lark_api"
	"option_get/log"
	"option_get/model"
	"strconv"
)

/*
	router
	OperationsResources.Post("/project_name")
	OperationsResources.Post("/resources_platform")
	OperationsResources.Post("/apply_type")
	OperationsResources.Post("/environment")
*/

type OpResourcesService struct {
}

func NewOpResourcesService() *OpResourcesService {
	return &OpResourcesService{}
}

func (or *OpResourcesService) GetProjectNameOptions(request *model.ApproveRequest) (*model.ApproveResponse, error) {
	appToken := request.LinkageParams.KV[constants.AppTokenKey]
	tableID := request.LinkageParams.KV[constants.TableIDKey]
	records, err := lark_api.GetBiTableRecords(appToken, tableID)
	if err != nil {
		return nil, err
	}
	response := model.NewApproveResponse()
	for i, item := range records.Data.Items {
		for key, field := range item.Fields {
			if values, ok := field.([]any); ok && key == constants.ProjectNameKey {
				for _, value := range values {
					var optionsID, optionsValue string
					optionsID = "options_" + strconv.Itoa(i) + "_id_" + ksuid.New().String()
					optionsValue = "@i18n@options_" + strconv.Itoa(i) + "_name_" + ksuid.New().String()
					response.Data.Result.Options = append(response.Data.Result.Options, &model.ApproveOption{
						IsDefault: false,
						ID:        optionsID,
						Value:     optionsValue,
					})
					for _, i18n := range response.Data.Result.I18NResources {
						switch i18n.Locale {
						case constants.ZhCNLocaleKey:
							i18n.Texts[optionsValue] = value.(map[string]interface{})[constants.TextKey].(string)
						case constants.EnUSLocalKey:
							//todo temporally unsupported
						default:
							return nil, errors.New("no such locale")
						}
					}
				}
			}
		}
	}
	return response, nil
}

func (or *OpResourcesService) GetResourcesPlatformOptions(request *model.ApproveRequest) (*model.ApproveResponse, error) {
	appToken := request.LinkageParams.KV[constants.AppTokenKey]
	tableToken := request.LinkageParams.KV[constants.TableIDKey]
	records, err := lark_api.GetBiTableFields(appToken, tableToken)
	if err != nil {
		return nil, err
	}
	response := model.NewApproveResponse()
	for _, item := range records.Data.Items {
		if item.FieldName == constants.ResourcesPlatformKey {
			for i, option := range item.Property.Options {
				var optionsID, optionsValue string
				optionsID = "options_" + strconv.Itoa(i) + "_id_" + ksuid.New().String()
				optionsValue = "@i18n@options_" + strconv.Itoa(i) + "_name_" + ksuid.New().String()
				response.Data.Result.Options = append(response.Data.Result.Options, &model.ApproveOption{
					IsDefault: false,
					ID:        optionsID,
					Value:     optionsValue,
				})
				for _, i18n := range response.Data.Result.I18NResources {
					switch i18n.Locale {
					case constants.ZhCNLocaleKey:
						i18n.Texts[optionsValue] = option.Name
					case constants.EnUSLocalKey:
						//todo temporally unsupported
					default:
						log.Debugf("no such locale \n")
					}
				}
			}
		}
	}
	return response, nil
}

// GetApplyTypeOptions todo
func (or *OpResourcesService) GetApplyTypeOptions(request *model.ApproveRequest) (*model.ApproveResponse, error) {
	appToken := request.LinkageParams.KV[constants.AppTokenKey]
	tableToken := request.LinkageParams.KV[constants.TableIDKey]
	records, err := lark_api.GetBiTableFields(appToken, tableToken)
	if err != nil {
		return nil, err
	}
	response := model.NewApproveResponse()
	for _, item := range records.Data.Items {
		log.Debugf("%#v \n", item)
		if item.FieldName == constants.ResourcesTypeKey {
			for i, option := range item.Property.Options {
				var optionsID, optionsValue string
				optionsID = "options_" + strconv.Itoa(i) + "_id_" + ksuid.New().String()
				optionsValue = "@i18n@options_" + strconv.Itoa(i) + "_name_" + ksuid.New().String()
				response.Data.Result.Options = append(response.Data.Result.Options, &model.ApproveOption{
					IsDefault: false,
					ID:        optionsID,
					Value:     optionsValue,
				})
				for _, i18n := range response.Data.Result.I18NResources {
					switch i18n.Locale {
					case constants.ZhCNLocaleKey:
						i18n.Texts[optionsValue] = option.Name
					case constants.EnUSLocalKey:
						//todo temporally unsupported
					default:
						log.Debugf("no such locale \n")
					}
				}
			}
		}
	}
	return response, nil
}

func (or *OpResourcesService) GetEnvOptions(request *model.ApproveRequest) (*model.ApproveResponse, error) {
	appToken := request.LinkageParams.KV[constants.AppTokenKey]
	tableToken := request.LinkageParams.KV[constants.TableIDKey]
	records, err := lark_api.GetBiTableFields(appToken, tableToken)
	if err != nil {
		return nil, err
	}
	response := model.NewApproveResponse()
	for _, item := range records.Data.Items {
		if item.FieldName == constants.EnvironmentKey {
			for i, option := range item.Property.Options {
				var optionsID, optionsValue string
				optionsID = "options_" + strconv.Itoa(i) + "_id_" + ksuid.New().String()
				optionsValue = "@i18n@options_" + strconv.Itoa(i) + "_name_" + ksuid.New().String()
				response.Data.Result.Options = append(response.Data.Result.Options, &model.ApproveOption{
					IsDefault: false,
					ID:        optionsID,
					Value:     optionsValue,
				})
				for _, i18n := range response.Data.Result.I18NResources {
					switch i18n.Locale {
					case constants.ZhCNLocaleKey:
						i18n.Texts[optionsValue] = option.Name
					case constants.EnUSLocalKey:
						//todo temporally unsupported
					default:
						log.Debugf("no such locale \n")
					}
				}
			}
		}
	}
	return response, nil
}
