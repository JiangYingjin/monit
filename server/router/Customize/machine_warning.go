package Customize

import (
	"github.com/gin-gonic/gin"
	"my-server/api/v1"
	"my-server/middleware"
)

type MachineWarningRouter struct {
}

// InitMachineWarningRouter 初始化 机器告警 路由信息
func (s *MachineWarningRouter) InitMachineWarningRouter(Router *gin.RouterGroup) {
	machineWarningRouter := Router.Group("machineWarning").Use(middleware.OperationRecord())
	machineWarningRouterWithoutRecord := Router.Group("machineWarning")
	var machineWarningApi = v1.ApiGroupApp.CustomizeApiGroup.MachineWarningApi
	{
		machineWarningRouter.POST("createMachineWarning", machineWarningApi.CreateMachineWarning)             // 新建机器告警
		machineWarningRouter.DELETE("deleteMachineWarning", machineWarningApi.DeleteMachineWarning)           // 删除机器告警
		machineWarningRouter.DELETE("deleteMachineWarningByIds", machineWarningApi.DeleteMachineWarningByIds) // 批量删除机器告警
		machineWarningRouter.PUT("updateMachineWarning", machineWarningApi.UpdateMachineWarning)              // 更新机器告警
	}
	{
		machineWarningRouterWithoutRecord.GET("findMachineWarning", machineWarningApi.FindMachineWarning)       // 根据ID获取机器告警
		machineWarningRouterWithoutRecord.GET("getMachineWarningList", machineWarningApi.GetMachineWarningList) // 获取机器告警列表
	}
}
