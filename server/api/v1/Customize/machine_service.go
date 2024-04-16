package Customize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Customize"
	CustomizeReq "github.com/flipped-aurora/gin-vue-admin/server/model/Customize/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MachineServiceApi struct {
}

var machineServiceService = service.ServiceGroupApp.CustomizeServiceGroup.MachineServiceService

// CreateMachineService 创建数据类型
// @Tags MachineService
// @Summary 创建数据类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Customize.MachineService true "创建数据类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /machineService/createMachineService [post]
func (machineServiceApi *MachineServiceApi) CreateMachineService(c *gin.Context) {
	var machineService Customize.MachineService
	err := c.ShouldBindJSON(&machineService)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	machineService.CreatedBy = utils.GetUserID(c)

	if err := machineServiceService.CreateMachineService(&machineService); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMachineService 删除数据类型
// @Tags MachineService
// @Summary 删除数据类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Customize.MachineService true "删除数据类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /machineService/deleteMachineService [delete]
func (machineServiceApi *MachineServiceApi) DeleteMachineService(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := machineServiceService.DeleteMachineService(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMachineServiceByIds 批量删除数据类型
// @Tags MachineService
// @Summary 批量删除数据类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /machineService/deleteMachineServiceByIds [delete]
func (machineServiceApi *MachineServiceApi) DeleteMachineServiceByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := machineServiceService.DeleteMachineServiceByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMachineService 更新数据类型
// @Tags MachineService
// @Summary 更新数据类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Customize.MachineService true "更新数据类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /machineService/updateMachineService [put]
func (machineServiceApi *MachineServiceApi) UpdateMachineService(c *gin.Context) {
	var machineService Customize.MachineService
	err := c.ShouldBindJSON(&machineService)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	machineService.UpdatedBy = utils.GetUserID(c)

	var oldMachineService Customize.MachineService
	err = global.GVA_DB.Where("machine_i_d = ?", machineService.MachineID).First(&oldMachineService).Error
	if err != nil {
		response.FailWithMessage("该数据不存在", c)
		return
	}
	oldMachineService.Services = machineService.Services

	if err = machineServiceService.UpdateMachineService(machineService); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMachineService 用id查询数据类型
// @Tags MachineService
// @Summary 用id查询数据类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Customize.MachineService true "用id查询数据类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /machineService/findMachineService [get]
func (machineServiceApi *MachineServiceApi) FindMachineService(c *gin.Context) {
	ID := c.Query("ID")
	if remachineService, err := machineServiceService.GetMachineService(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remachineService": remachineService}, c)
	}
}

// GetMachineServiceList 分页获取数据类型列表
// @Tags MachineService
// @Summary 分页获取数据类型列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query CustomizeReq.MachineServiceSearch true "分页获取数据类型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /machineService/getMachineServiceList [get]
func (machineServiceApi *MachineServiceApi) GetMachineServiceList(c *gin.Context) {
	var pageInfo CustomizeReq.MachineServiceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := machineServiceService.GetMachineServiceInfoList(pageInfo); err != nil {
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
