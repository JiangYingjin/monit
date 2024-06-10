package Customize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MachineWarningLogRouter struct {
}

// InitMachineWarningLogRouter 初始化 machineWarningLog表 路由信息
func (s *MachineWarningLogRouter) InitMachineWarningLogRouter(Router *gin.RouterGroup) {
	machineWarningLogRouter := Router.Group("machineWarningLog").Use(middleware.OperationRecord())
	machineWarningLogRouterWithoutRecord := Router.Group("machineWarningLog")
	var machineWarningLogApi = v1.ApiGroupApp.CustomizeApiGroup.MachineWarningLogApi
	{
		machineWarningLogRouter.POST("createMachineWarningLog", machineWarningLogApi.CreateMachineWarningLog)             // 新建machineWarningLog表
		machineWarningLogRouter.DELETE("deleteMachineWarningLog", machineWarningLogApi.DeleteMachineWarningLog)           // 删除machineWarningLog表
		machineWarningLogRouter.DELETE("deleteMachineWarningLogByIds", machineWarningLogApi.DeleteMachineWarningLogByIds) // 批量删除machineWarningLog表
		machineWarningLogRouter.PUT("updateMachineWarningLog", machineWarningLogApi.UpdateMachineWarningLog)              // 更新machineWarningLog表
	}
	{
		machineWarningLogRouterWithoutRecord.GET("findMachineWarningLog", machineWarningLogApi.FindMachineWarningLog)       // 根据ID获取machineWarningLog表
		machineWarningLogRouterWithoutRecord.GET("getMachineWarningLogList", machineWarningLogApi.GetMachineWarningLogList) // 获取machineWarningLog表列表
	}
}
