package Customize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Customize"
	CustomizeReq "github.com/flipped-aurora/gin-vue-admin/server/model/Customize/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MachineWarningLogApi struct {
}

var machineWarningLogService = service.ServiceGroupApp.CustomizeServiceGroup.MachineWarningLogService

// CreateMachineWarningLog 创建machineWarningLog表
// @Tags MachineWarningLog
// @Summary 创建machineWarningLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Customize.MachineWarningLog true "创建machineWarningLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /machineWarningLog/createMachineWarningLog [post]
func (machineWarningLogApi *MachineWarningLogApi) CreateMachineWarningLog(c *gin.Context) {
	var machineWarningLog Customize.MachineWarningLog
	err := c.ShouldBindJSON(&machineWarningLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := machineWarningLogService.CreateMachineWarningLog(&machineWarningLog); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMachineWarningLog 删除machineWarningLog表
// @Tags MachineWarningLog
// @Summary 删除machineWarningLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Customize.MachineWarningLog true "删除machineWarningLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /machineWarningLog/deleteMachineWarningLog [delete]
func (machineWarningLogApi *MachineWarningLogApi) DeleteMachineWarningLog(c *gin.Context) {
	userId := c.Query("userId")
	if err := machineWarningLogService.DeleteMachineWarningLog(userId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMachineWarningLogByIds 批量删除machineWarningLog表
// @Tags MachineWarningLog
// @Summary 批量删除machineWarningLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /machineWarningLog/deleteMachineWarningLogByIds [delete]
func (machineWarningLogApi *MachineWarningLogApi) DeleteMachineWarningLogByIds(c *gin.Context) {
	userIds := c.QueryArray("userIds[]")
	if err := machineWarningLogService.DeleteMachineWarningLogByIds(userIds); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMachineWarningLog 更新machineWarningLog表
// @Tags MachineWarningLog
// @Summary 更新machineWarningLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Customize.MachineWarningLog true "更新machineWarningLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /machineWarningLog/updateMachineWarningLog [put]
func (machineWarningLogApi *MachineWarningLogApi) UpdateMachineWarningLog(c *gin.Context) {
	var machineWarningLog Customize.MachineWarningLog
	err := c.ShouldBindJSON(&machineWarningLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := machineWarningLogService.UpdateMachineWarningLog(machineWarningLog); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMachineWarningLog 用id查询machineWarningLog表
// @Tags MachineWarningLog
// @Summary 用id查询machineWarningLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Customize.MachineWarningLog true "用id查询machineWarningLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /machineWarningLog/findMachineWarningLog [get]
func (machineWarningLogApi *MachineWarningLogApi) FindMachineWarningLog(c *gin.Context) {
	userId := c.Query("userId")
	if remachineWarningLog, err := machineWarningLogService.GetMachineWarningLog(userId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remachineWarningLog": remachineWarningLog}, c)
	}
}

// GetMachineWarningLogList 分页获取machineWarningLog表列表
// @Tags MachineWarningLog
// @Summary 分页获取machineWarningLog表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query CustomizeReq.MachineWarningLogSearch true "分页获取machineWarningLog表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /machineWarningLog/getMachineWarningLogList [get]
func (machineWarningLogApi *MachineWarningLogApi) GetMachineWarningLogList(c *gin.Context) {
	var pageInfo CustomizeReq.MachineWarningLogSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := machineWarningLogService.GetMachineWarningLogInfoList(pageInfo); err != nil {
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
