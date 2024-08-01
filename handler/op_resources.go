package handler

import (
	"github.com/kataras/iris/v12"
	"option_get/constants"
	"option_get/model"
	"option_get/service"
)

/*
	OperationsResources.Post("/project_name")
	OperationsResources.Post("/resources_platform")
	OperationsResources.Post("/apply_type")
	OperationsResources.Post("/environment")
*/

func GetOpResourcesProjectNameOptions(ctx iris.Context) {
	tableID := ctx.URLParam("table_id")
	appToken := ctx.URLParam("app_token")
	opResourcesService := service.NewOpResourcesService()
	approveRequest := model.NewApproveRequest()
	approveRequest.LinkageParams.KV[constants.AppTokenKey] = appToken
	approveRequest.LinkageParams.KV[constants.TableIDKey] = tableID
	options, err := opResourcesService.GetProjectNameOptions(approveRequest)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		responseFailed(ctx, err)
		return
	}
	ctx.StatusCode(iris.StatusOK)
	responseSuccessful(ctx, 0, "success", options.Data)
}

func GetOpResourcesPlatformOptions(ctx iris.Context) {
	tableID := ctx.URLParam("table_id")
	appToken := ctx.URLParam("app_token")
	opResourcesService := service.NewOpResourcesService()
	approveRequest := model.NewApproveRequest()
	approveRequest.LinkageParams.KV[constants.AppTokenKey] = appToken
	approveRequest.LinkageParams.KV[constants.TableIDKey] = tableID
	options, err := opResourcesService.GetResourcesPlatformOptions(approveRequest)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		responseFailed(ctx, err)
		return
	}
	ctx.StatusCode(iris.StatusOK)
	responseSuccessful(ctx, 0, "success", options.Data)
}

func GetOpResourcesApplyTypeOptions(ctx iris.Context) {
	tableID := ctx.URLParam("table_id")
	appToken := ctx.URLParam("app_token")
	opResourcesService := service.NewOpResourcesService()
	approveRequest := model.NewApproveRequest()
	approveRequest.LinkageParams.KV[constants.AppTokenKey] = appToken
	approveRequest.LinkageParams.KV[constants.TableIDKey] = tableID
	options, err := opResourcesService.GetApplyTypeOptions(approveRequest)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		responseFailed(ctx, err)
		return
	}
	ctx.StatusCode(iris.StatusOK)
	responseSuccessful(ctx, 0, "success", options.Data)
}

func GetOpResourcesEnvOptions(ctx iris.Context) {
	tableID := ctx.URLParam("table_id")
	appToken := ctx.URLParam("app_token")
	opResourcesService := service.NewOpResourcesService()
	approveRequest := model.NewApproveRequest()
	approveRequest.LinkageParams.KV[constants.AppTokenKey] = appToken
	approveRequest.LinkageParams.KV[constants.TableIDKey] = tableID
	options, err := opResourcesService.GetEnvOptions(approveRequest)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		responseFailed(ctx, err)
		return
	}
	ctx.StatusCode(iris.StatusOK)
	responseSuccessful(ctx, 0, "success", options.Data)
}
