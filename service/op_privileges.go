package service

import "option_get/model"

type OpPrivilegesService struct{}

func NewOpPrivilegesService() *OpPrivilegesService {
	return &OpPrivilegesService{}
}

func (op *OpPrivilegesService) GetOpPrivilegeSystemNameOptions(request *model.ApproveRequest) (*model.ApproveResponse, error) {
	response := model.NewApproveResponse()

	return response, nil
}
