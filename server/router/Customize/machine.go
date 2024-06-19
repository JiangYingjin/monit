package Customize

import (
	"github.com/gin-gonic/gin"
	"my-server/api/v1"
	"my-server/middleware"
)

type MachineRouter struct {
}

// InitMachineRouter 初始化 Machine 路由信息
func (s *MachineRouter) InitMachineRouter(Router *gin.RouterGroup) {
	machineRouter := Router.Group("machine").Use(middleware.OperationRecord())
	machineRouterWithoutRecord := Router.Group("machine")
	var machineApi = v1.ApiGroupApp.CustomizeApiGroup.MachineApi
	var myMachineApi = v1.ApiGroupApp.CustomizeApiGroup.MyMachineApi
	{
		machineRouter.POST("createMachine", machineApi.CreateMachine)             // 新建Machine
		machineRouter.DELETE("deleteMachine", machineApi.DeleteMachine)           // 删除Machine
		machineRouter.DELETE("deleteMachineByIds", machineApi.DeleteMachineByIds) // 批量删除Machine
		machineRouter.PUT("updateMachine", machineApi.UpdateMachine)              // 更新Machine

		machineRouter.POST("setMachineService", myMachineApi.SetMachineService) // 设置服务监听状态
	}
	{
		machineRouterWithoutRecord.GET("findMachine", machineApi.FindMachine)       // 根据ID获取Machine
		machineRouterWithoutRecord.GET("getMachineList", machineApi.GetMachineList) // 获取Machine列表

	}
}
