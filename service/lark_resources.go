package service

import (
	"errors"
	"github.com/segmentio/ksuid"
	"option_get/constants"
	"option_get/lark_api"
	"option_get/model"
	"option_get/utils/common"
	"strconv"
)

type LarkResourcesService struct{}

func NewLarkResourcesService() *LarkResourcesService {
	return &LarkResourcesService{}
}

func (lr *LarkResourcesService) GetApplySystemNameOptions(request *model.ApproveRequest) (*model.ApproveResponse, error) {
	appToken := request.LinkageParams.KV[constants.AppTokenKey]
	tableID := request.LinkageParams.KV[constants.TableIDKey]
	records, err := lark_api.GetBiTableRecords(appToken, tableID)
	if err != nil {
		return nil, err
	}
	response := model.NewApproveResponse()

	//寻找到对应的唯一的父记录的记录号
	targetRecordID := records.GetRecordIDByName(constants.LarkResourcesApplyNameKey)
	for i, item := range records.Data.Items {
		fieldMap, ok1 := item.Fields[constants.FatherRecordKey].(map[string]interface{})
		recordsSlices, ok2 := fieldMap[constants.LinkRecordIdsKey].([]interface{})
		if ok1 {
			if !ok2 || recordsSlices == nil || !common.InterfaceSliceContain(recordsSlices, targetRecordID) {
				continue
			}
		}
		if values, ok := item.Fields[constants.OptionsKey].([]any); ok {
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
	return response, nil
}
