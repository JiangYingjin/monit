package Customize

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type MyMachineRouter struct {
}

// InitMyMachineRouter 初始化 MyMachine 路由信息
func (s *MachineRouter) InitMyMachineRouter(Router *gin.RouterGroup) {
	var machineApi = v1.ApiGroupApp.CustomizeApiGroup.MyMachineApi

	routerGroup := Router.Group("machine")
	routerGroup.POST("login", machineApi.Login) // 根据ID获取Machine
}
