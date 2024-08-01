package handler

import (
	"github.com/kataras/iris/v12"
	"option_get/constants"
	"option_get/model"
	"option_get/service"
)

func GetNetResourcesApplyNetworkNameOptions(ctx iris.Context) {
	tableID := ctx.URLParam("table_id")
	appToken := ctx.URLParam("app_token")
	netResourcesService := service.NewNetResourcesService()
	approveRequest := model.NewApproveRequest()
	approveRequest.LinkageParams.KV[constants.AppTokenKey] = appToken
	approveRequest.LinkageParams.KV[constants.TableIDKey] = tableID
	options, err := netResourcesService.GetApplyNetworkNameOptions(approveRequest)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		responseFailed(ctx, err)
		return
	}
	ctx.StatusCode(iris.StatusOK)
	responseSuccessful(ctx, 0, "success", options.Data)
}
