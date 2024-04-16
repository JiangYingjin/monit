package Customize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ServiceTemplateRouter struct {
}

// InitServiceTemplateRouter 初始化 命令模板 路由信息
func (s *ServiceTemplateRouter) InitServiceTemplateRouter(Router *gin.RouterGroup) {
	serviceTemplateRouter := Router.Group("serviceTemplate").Use(middleware.OperationRecord())
	serviceTemplateRouterWithoutRecord := Router.Group("serviceTemplate")
	var serviceTemplateApi = v1.ApiGroupApp.CustomizeApiGroup.ServiceTemplateApi
	{
		serviceTemplateRouter.POST("createServiceTemplate", serviceTemplateApi.CreateServiceTemplate)             // 新建命令模板
		serviceTemplateRouter.DELETE("deleteServiceTemplate", serviceTemplateApi.DeleteServiceTemplate)           // 删除命令模板
		serviceTemplateRouter.DELETE("deleteServiceTemplateByIds", serviceTemplateApi.DeleteServiceTemplateByIds) // 批量删除命令模板
		serviceTemplateRouter.PUT("updateServiceTemplate", serviceTemplateApi.UpdateServiceTemplate)              // 更新命令模板
	}
	{
		serviceTemplateRouterWithoutRecord.GET("findServiceTemplate", serviceTemplateApi.FindServiceTemplate)       // 根据ID获取命令模板
		serviceTemplateRouterWithoutRecord.GET("getServiceTemplateList", serviceTemplateApi.GetServiceTemplateList) // 获取命令模板列表
	}
}
