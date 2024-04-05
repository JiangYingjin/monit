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

type DataApi struct {
}

var dataService = service.ServiceGroupApp.CustomizeServiceGroup.DataService

// CreateData 创建Data
// @Tags Data
// @Summary 创建Data
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Customize.Data true "创建Data"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /data/createData [post]
//
//	{
//		"dataTypeID": 1,
//		"machineID": 1,
//		"value": 1.0
//	}
func (dataApi *DataApi) CreateData(c *gin.Context) {
	var data Customize.Data
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data.CreatedBy = uint(*data.MachineID)

	if err := dataService.CreateData(&data); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteData 删除Data
// @Tags Data
// @Summary 删除Data
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Customize.Data true "删除Data"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /data/deleteData [delete]
func (dataApi *DataApi) DeleteData(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := dataService.DeleteData(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDataByIds 批量删除Data
// @Tags Data
// @Summary 批量删除Data
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /data/deleteDataByIds [delete]
func (dataApi *DataApi) DeleteDataByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := dataService.DeleteDataByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateData 更新Data
// @Tags Data
// @Summary 更新Data
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Customize.Data true "更新Data"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /data/updateData [put]
func (dataApi *DataApi) UpdateData(c *gin.Context) {
	var data Customize.Data
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data.UpdatedBy = utils.GetUserID(c)

	if err := dataService.UpdateData(data); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindData 用id查询Data
// @Tags Data
// @Summary 用id查询Data
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Customize.Data true "用id查询Data"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /data/findData [get]
func (dataApi *DataApi) FindData(c *gin.Context) {
	ID := c.Query("ID")
	if redata, err := dataService.GetData(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redata": redata}, c)
	}
}

// GetDataList 分页获取Data列表
// @Tags Data
// @Summary 分页获取Data列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query CustomizeReq.DataSearch true "分页获取Data列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /data/getDataList [get]
func (dataApi *DataApi) GetDataList(c *gin.Context) {
	var pageInfo CustomizeReq.DataSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := dataService.GetDataInfoList(pageInfo); err != nil {
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
