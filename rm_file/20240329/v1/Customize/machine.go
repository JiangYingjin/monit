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

type MachineApi struct {
}

var machineService = service.ServiceGroupApp.CustomizeServiceGroup.MachineService

// CreateMachine 创建Machine
// @Tags Machine
// @Summary 创建Machine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Customize.Machine true "创建Machine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /machine/createMachine [post]
func (machineApi *MachineApi) CreateMachine(c *gin.Context) {
	var machine Customize.Machine
	err := c.ShouldBindJSON(&machine)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	machine.CreatedBy = utils.GetUserID(c)

	if err := machineService.CreateMachine(&machine); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMachine 删除Machine
// @Tags Machine
// @Summary 删除Machine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Customize.Machine true "删除Machine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /machine/deleteMachine [delete]
func (machineApi *MachineApi) DeleteMachine(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := machineService.DeleteMachine(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMachineByIds 批量删除Machine
// @Tags Machine
// @Summary 批量删除Machine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /machine/deleteMachineByIds [delete]
func (machineApi *MachineApi) DeleteMachineByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := machineService.DeleteMachineByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMachine 更新Machine
// @Tags Machine
// @Summary 更新Machine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Customize.Machine true "更新Machine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /machine/updateMachine [put]
func (machineApi *MachineApi) UpdateMachine(c *gin.Context) {
	var machine Customize.Machine
	err := c.ShouldBindJSON(&machine)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	machine.UpdatedBy = utils.GetUserID(c)

	if err := machineService.UpdateMachine(machine); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMachine 用id查询Machine
// @Tags Machine
// @Summary 用id查询Machine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Customize.Machine true "用id查询Machine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /machine/findMachine [get]
func (machineApi *MachineApi) FindMachine(c *gin.Context) {
	ID := c.Query("ID")
	if remachine, err := machineService.GetMachine(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remachine": remachine}, c)
	}
}

// GetMachineList 分页获取Machine列表
// @Tags Machine
// @Summary 分页获取Machine列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query CustomizeReq.MachineSearch true "分页获取Machine列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /machine/getMachineList [get]
func (machineApi *MachineApi) GetMachineList(c *gin.Context) {
	var pageInfo CustomizeReq.MachineSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := machineService.GetMachineInfoList(pageInfo); err != nil {
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

type MachineLoginReq struct {
	machineID       string
	machinePassword string
}

// MachineLogin 分页获取Machine列表
// @Tags Machine
// @Summary 分页获取Machine列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query CustomizeReq.MachineSearch true "分页获取Machine列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /machine/getMachineList [get]
func (machineApi *MachineApi) MachineLogin(c *gin.Context) {
	response.OK(c)
}
