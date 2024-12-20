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

type MachineWarningApi struct {
}

var machineWarningService = service.ServiceGroupApp.CustomizeServiceGroup.MachineWarningService

// CreateMachineWarning 创建机器告警
// @Tags MachineWarning
// @Summary 创建机器告警
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Customize.MachineWarning true "创建机器告警"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /machineWarning/createMachineWarning [post]
func (machineWarningApi *MachineWarningApi) CreateMachineWarning(c *gin.Context) {
	var machineWarning Customize.MachineWarning
	err := c.ShouldBindJSON(&machineWarning)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	machineWarning.CreatedBy = utils.GetUserID(c)

	if err := machineWarningService.CreateMachineWarning(&machineWarning); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMachineWarning 删除机器告警
// @Tags MachineWarning
// @Summary 删除机器告警
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Customize.MachineWarning true "删除机器告警"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /machineWarning/deleteMachineWarning [delete]
func (machineWarningApi *MachineWarningApi) DeleteMachineWarning(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := machineWarningService.DeleteMachineWarning(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMachineWarningByIds 批量删除机器告警
// @Tags MachineWarning
// @Summary 批量删除机器告警
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /machineWarning/deleteMachineWarningByIds [delete]
func (machineWarningApi *MachineWarningApi) DeleteMachineWarningByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := machineWarningService.DeleteMachineWarningByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMachineWarning 更新机器告警
// @Tags MachineWarning
// @Summary 更新机器告警
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Customize.MachineWarning true "更新机器告警"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /machineWarning/updateMachineWarning [put]
func (machineWarningApi *MachineWarningApi) UpdateMachineWarning(c *gin.Context) {
	var machineWarning Customize.MachineWarning
	err := c.ShouldBindJSON(&machineWarning)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	machineWarning.UpdatedBy = utils.GetUserID(c)

	if err := machineWarningService.UpdateMachineWarning(machineWarning); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMachineWarning 用id查询机器告警
// @Tags MachineWarning
// @Summary 用id查询机器告警
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Customize.MachineWarning true "用id查询机器告警"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /machineWarning/findMachineWarning [get]
func (machineWarningApi *MachineWarningApi) FindMachineWarning(c *gin.Context) {
	ID := c.Query("ID")
	if remachineWarning, err := machineWarningService.GetMachineWarning(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remachineWarning": remachineWarning}, c)
	}
}

// GetMachineWarningList 分页获取机器告警列表
// @Tags MachineWarning
// @Summary 分页获取机器告警列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query CustomizeReq.MachineWarningSearch true "分页获取机器告警列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /machineWarning/getMachineWarningList [get]
func (machineWarningApi *MachineWarningApi) GetMachineWarningList(c *gin.Context) {
	var pageInfo CustomizeReq.MachineWarningSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := machineWarningService.GetMachineWarningInfoList(pageInfo); err != nil {
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
