package Customize

import (
	"github.com/gin-gonic/gin"
	v1 "my-server/api/v1"
	"my-server/api/v1/Customize"
)

type MyMachineRouter struct {
}

// InitMyMachineRouter 初始化 MyMachine 路由信息
func (s *MachineRouter) InitMyMachineLoginRouter(Router *gin.RouterGroup) {
	var machineApi = v1.ApiGroupApp.CustomizeApiGroup.MyMachineApi

	routerGroup := Router.Group("machine")
	routerGroup.POST("login", machineApi.MachineLogin) // 根据ID获取Machine
}

func (s *MachineRouter) InitMyMachineRouter(Router *gin.RouterGroup) {
	//var machineApi = v1.ApiGroupApp.CustomizeApiGroup.MyMachineApi
	var dataApi = v1.ApiGroupApp.CustomizeApiGroup.DataApi

	routerGroup := Router.Group("machine")
	routerGroup.POST("uploadData", dataApi.CreateData)
	routerGroup.POST("uploadDataMulti", dataApi.CreateDataMulti)

	var machineServiceApi Customize.MyMachineApi
	routerGroup.POST("updateMachineService", machineServiceApi.UpdateMachineService) // 更新数据类型
}
