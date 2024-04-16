package Customize

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Customize"
	"github.com/gin-gonic/gin"
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

	var machineServiceApi Customize.MachineServiceApi
	routerGroup.POST("createMachineService", machineServiceApi.CreateMachineService) // 新建数据类型
	routerGroup.PUT("updateMachineService", machineServiceApi.UpdateMachineService)  // 更新数据类型
}
