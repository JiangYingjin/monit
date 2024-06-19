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

type DataTypeApi struct {
}

var dataTypeService = service.ServiceGroupApp.CustomizeServiceGroup.DataTypeService

// CreateDataType 创建DataType
// @Tags DataType
// @Summary 创建DataType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Customize.DataType true "创建DataType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dataType/createDataType [post]
func (dataTypeApi *DataTypeApi) CreateDataType(c *gin.Context) {
	var dataType Customize.DataType
	err := c.ShouldBindJSON(&dataType)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	dataType.CreatedBy = utils.GetUserID(c)

	if err := dataTypeService.CreateDataType(&dataType); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDataType 删除DataType
// @Tags DataType
// @Summary 删除DataType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Customize.DataType true "删除DataType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dataType/deleteDataType [delete]
func (dataTypeApi *DataTypeApi) DeleteDataType(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := dataTypeService.DeleteDataType(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDataTypeByIds 批量删除DataType
// @Tags DataType
// @Summary 批量删除DataType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dataType/deleteDataTypeByIds [delete]
func (dataTypeApi *DataTypeApi) DeleteDataTypeByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := dataTypeService.DeleteDataTypeByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDataType 更新DataType
// @Tags DataType
// @Summary 更新DataType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Customize.DataType true "更新DataType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dataType/updateDataType [put]
func (dataTypeApi *DataTypeApi) UpdateDataType(c *gin.Context) {
	var dataType Customize.DataType
	err := c.ShouldBindJSON(&dataType)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	dataType.UpdatedBy = utils.GetUserID(c)

	if err := dataTypeService.UpdateDataType(dataType); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDataType 用id查询DataType
// @Tags DataType
// @Summary 用id查询DataType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Customize.DataType true "用id查询DataType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dataType/findDataType [get]
func (dataTypeApi *DataTypeApi) FindDataType(c *gin.Context) {
	ID := c.Query("ID")
	if redataType, err := dataTypeService.GetDataType(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redataType": redataType}, c)
	}
}

// GetDataTypeList 分页获取DataType列表
// @Tags DataType
// @Summary 分页获取DataType列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query CustomizeReq.DataTypeSearch true "分页获取DataType列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dataType/getDataTypeList [get]
func (dataTypeApi *DataTypeApi) GetDataTypeList(c *gin.Context) {
	var pageInfo CustomizeReq.DataTypeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := dataTypeService.GetDataTypeInfoList(pageInfo); err != nil {
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
