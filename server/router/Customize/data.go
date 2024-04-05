package Customize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DataRouter struct {
}

// InitDataRouter 初始化 Data 路由信息
func (s *DataRouter) InitDataRouter(Router *gin.RouterGroup) {
	dataRouter := Router.Group("data").Use(middleware.OperationRecord())
	dataRouterWithoutRecord := Router.Group("data")
	var dataApi = v1.ApiGroupApp.CustomizeApiGroup.DataApi
	var machineApi = v1.ApiGroupApp.CustomizeApiGroup.MyMachineApi
	{
		dataRouter.POST("createData", dataApi.CreateData)             // 新建Data
		dataRouter.DELETE("deleteData", dataApi.DeleteData)           // 删除Data
		dataRouter.DELETE("deleteDataByIds", dataApi.DeleteDataByIds) // 批量删除Data
		dataRouter.PUT("updateData", dataApi.UpdateData)              // 更新Data
	}
	{
		dataRouterWithoutRecord.GET("findData", dataApi.FindData)       // 根据ID获取Data
		dataRouterWithoutRecord.GET("getDataList", dataApi.GetDataList) // 获取Data列表
		dataRouterWithoutRecord.POST("getData", machineApi.GetData)     // 获取Data
	}
}
