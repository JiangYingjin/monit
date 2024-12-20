package Customize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MachineRouter struct {
}

// InitMachineRouter 初始化 Machine 路由信息
func (s *MachineRouter) InitMachineRouter(Router *gin.RouterGroup) {
	machineRouter := Router.Group("machine").Use(middleware.OperationRecord())
	machineRouterWithoutRecord := Router.Group("machine")
	var machineApi = v1.ApiGroupApp.CustomizeApiGroup.MachineApi
	{
		machineRouter.POST("createMachine", machineApi.CreateMachine)             // 新建Machine
		machineRouter.DELETE("deleteMachine", machineApi.DeleteMachine)           // 删除Machine
		machineRouter.DELETE("deleteMachineByIds", machineApi.DeleteMachineByIds) // 批量删除Machine
		machineRouter.PUT("updateMachine", machineApi.UpdateMachine)              // 更新Machine
	}
	{
		machineRouterWithoutRecord.GET("findMachine", machineApi.FindMachine)       // 根据ID获取Machine
		machineRouterWithoutRecord.GET("getMachineList", machineApi.GetMachineList) // 获取Machine列表
	}
}
