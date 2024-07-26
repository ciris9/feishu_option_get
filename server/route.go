package server

import (
	"github.com/kataras/iris/v12"
	"option_get/handler"
)

/*
todo
has relation
运维资源申请 项目名称 部署平台

todo
has no relation
运维资源申请 申请平台 所属环境 申请类型
运维权限申请 系统名称
网络资源申请 申请系统名称 设备类型
飞书资源申请 申请系统名称
*/

func InitRoute(app *iris.Application) {
	//todo 运维资源申请
	OperationsResources := app.Party("/op_resources_options")
	{
		OperationsResources.Post("/project_name", handler.GetOpResourcesProjectNameOptions)
		OperationsResources.Post("/resources_platform", handler.GetOpResourcesPlatformOptions)
		OperationsResources.Post("/apply_type", handler.GetOpResourcesApplyTypeOptions)
		OperationsResources.Post("/environment", handler.GetOpResourcesEnvOptions)
	}

	//todo 运维权限申请
	OperationsPrivileges := app.Party("/op_privileges_options")
	{
		OperationsPrivileges.Post("/system_name", handler.GetOpPrivilegeSystemNameOptions)
	}

	//todo 网络资源申请
	NetworkResources := app.Party("/net_resources_options")
	{
		NetworkResources.Post("/apply_network_name", handler.GetNetResourcesApplyNetworkNameOptions)
	}

	//todo 飞书资源申请
	LarkResources := app.Party("/lark_resources_options")
	{
		LarkResources.Post("/apply_system_name", handler.GetLarkResourcesApplySystemNameOptions)
	}
}
