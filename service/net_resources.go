package service

import "option_get/model"

type NetResourcesService struct {
}

func NewNetResourcesService() *NetResourcesService {
	return &NetResourcesService{}
}

func (nr *NetResourcesService) GetApplyNetworkNameOptions(request *model.ApproveRequest) (*model.ApproveResponse, error) {
	response := model.NewApproveResponse()

	return response, nil
}
