package Customize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DataTypeRouter struct {
}

// InitDataTypeRouter 初始化 DataType 路由信息
func (s *DataTypeRouter) InitDataTypeRouter(Router *gin.RouterGroup) {
	dataTypeRouter := Router.Group("dataType").Use(middleware.OperationRecord())
	dataTypeRouterWithoutRecord := Router.Group("dataType")
	var dataTypeApi = v1.ApiGroupApp.CustomizeApiGroup.DataTypeApi
	{
		dataTypeRouter.POST("createDataType", dataTypeApi.CreateDataType)             // 新建DataType
		dataTypeRouter.DELETE("deleteDataType", dataTypeApi.DeleteDataType)           // 删除DataType
		dataTypeRouter.DELETE("deleteDataTypeByIds", dataTypeApi.DeleteDataTypeByIds) // 批量删除DataType
		dataTypeRouter.PUT("updateDataType", dataTypeApi.UpdateDataType)              // 更新DataType
	}
	{
		dataTypeRouterWithoutRecord.GET("findDataType", dataTypeApi.FindDataType)       // 根据ID获取DataType
		dataTypeRouterWithoutRecord.GET("getDataTypeList", dataTypeApi.GetDataTypeList) // 获取DataType列表
	}
}
