package Customize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
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
