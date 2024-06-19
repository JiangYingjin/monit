package Customize

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"my-server/global"
	"my-server/model/Customize"
	CustomizeReq "my-server/model/Customize/request"
	"my-server/model/common/response"
	"my-server/service"
	"my-server/utils"
)

type ServiceTemplateApi struct {
}

var serviceTemplateService = service.ServiceGroupApp.CustomizeServiceGroup.ServiceTemplateService

// CreateServiceTemplate 创建命令模板
// @Tags ServiceTemplate
// @Summary 创建命令模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Customize.ServiceTemplate true "创建命令模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /serviceTemplate/createServiceTemplate [post]
func (serviceTemplateApi *ServiceTemplateApi) CreateServiceTemplate(c *gin.Context) {
	var serviceTemplate Customize.ServiceTemplate
	err := c.ShouldBindJSON(&serviceTemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	serviceTemplate.CreatedBy = utils.GetUserID(c)

	if err := serviceTemplateService.CreateServiceTemplate(&serviceTemplate); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteServiceTemplate 删除命令模板
// @Tags ServiceTemplate
// @Summary 删除命令模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Customize.ServiceTemplate true "删除命令模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /serviceTemplate/deleteServiceTemplate [delete]
func (serviceTemplateApi *ServiceTemplateApi) DeleteServiceTemplate(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := serviceTemplateService.DeleteServiceTemplate(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteServiceTemplateByIds 批量删除命令模板
// @Tags ServiceTemplate
// @Summary 批量删除命令模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /serviceTemplate/deleteServiceTemplateByIds [delete]
func (serviceTemplateApi *ServiceTemplateApi) DeleteServiceTemplateByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := serviceTemplateService.DeleteServiceTemplateByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateServiceTemplate 更新命令模板
// @Tags ServiceTemplate
// @Summary 更新命令模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Customize.ServiceTemplate true "更新命令模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /serviceTemplate/updateServiceTemplate [put]
func (serviceTemplateApi *ServiceTemplateApi) UpdateServiceTemplate(c *gin.Context) {
	var serviceTemplate Customize.ServiceTemplate
	err := c.ShouldBindJSON(&serviceTemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	serviceTemplate.UpdatedBy = utils.GetUserID(c)

	if err := serviceTemplateService.UpdateServiceTemplate(serviceTemplate); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindServiceTemplate 用id查询命令模板
// @Tags ServiceTemplate
// @Summary 用id查询命令模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Customize.ServiceTemplate true "用id查询命令模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /serviceTemplate/findServiceTemplate [get]
func (serviceTemplateApi *ServiceTemplateApi) FindServiceTemplate(c *gin.Context) {
	ID := c.Query("ID")
	if reserviceTemplate, err := serviceTemplateService.GetServiceTemplate(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reserviceTemplate": reserviceTemplate}, c)
	}
}

// GetServiceTemplateList 分页获取命令模板列表
// @Tags ServiceTemplate
// @Summary 分页获取命令模板列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query CustomizeReq.ServiceTemplateSearch true "分页获取命令模板列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /serviceTemplate/getServiceTemplateList [get]
func (serviceTemplateApi *ServiceTemplateApi) GetServiceTemplateList(c *gin.Context) {
	var pageInfo CustomizeReq.ServiceTemplateSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := serviceTemplateService.GetServiceTemplateInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
