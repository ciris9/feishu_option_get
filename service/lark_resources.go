package service

import "option_get/model"

type LarkResourcesService struct{}

func (lr *LarkResourcesService) GetApplySystemNameOptions(request *model.ApproveRequest) (*model.ApproveResponse, error) {
	response := model.NewApproveResponse()

	return response, nil
}
