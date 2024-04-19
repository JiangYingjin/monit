package Customize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MachineServiceRouter struct {
}

// InitMachineServiceRouter 初始化 数据类型 路由信息
func (s *MachineServiceRouter) InitMachineServiceRouter(Router *gin.RouterGroup) {
	machineServiceRouter := Router.Group("machineService").Use(middleware.OperationRecord())
	machineServiceRouterWithoutRecord := Router.Group("machineService")
	var machineServiceApi = v1.ApiGroupApp.CustomizeApiGroup.MachineServiceApi
	{
		machineServiceRouter.POST("createMachineService", machineServiceApi.CreateMachineService)             // 新建数据类型
		machineServiceRouter.DELETE("deleteMachineService", machineServiceApi.DeleteMachineService)           // 删除数据类型
		machineServiceRouter.DELETE("deleteMachineServiceByIds", machineServiceApi.DeleteMachineServiceByIds) // 批量删除数据类型
		machineServiceRouter.PUT("updateMachineService", machineServiceApi.UpdateMachineService)              // 更新数据类型
	}
	{
		machineServiceRouterWithoutRecord.GET("findMachineService", machineServiceApi.FindMachineService)       // 根据ID获取数据类型
		machineServiceRouterWithoutRecord.GET("getMachineServiceList", machineServiceApi.GetMachineServiceList) // 获取数据类型列表
	}
}
