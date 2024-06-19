package Customize

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"my-server/global"
	"my-server/model/Customize"
	CustomizeReq "my-server/model/Customize/request"
	"my-server/model/common/response"
	"my-server/service"
	Customize2 "my-server/service/Customize"
	"my-server/utils"
	"time"
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
type CreateDataReq struct {
	DataTypeID *int     `json:"dataTypeID" form:"dataTypeID" gorm:"column:data_type_i_d;comment:;" binding:"required"` //数据类型
	Value      *float64 `json:"value" form:"value" gorm:"column:value;comment:;" binding:"required"`                   //值
	MachineID  *int     `json:"machineID" form:"machineID" gorm:"column:machine_i_d;comment:;" binding:"required"`     //机器ID
	CreatedAt  string   `json:"created_at"`                                                                            // 创建时间
	UpdatedAt  string   `json:"updated_at"`                                                                            // 更新时间
}

func (req *CreateDataReq) ToData() (Customize.Data, error) {
	createdAt, err := time.Parse(time.RFC3339, Customize2.ConvertTimestamp(req.CreatedAt))
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		return Customize.Data{}, err
	}
	return Customize.Data{
		DataTypeID: req.DataTypeID,
		Value:      req.Value,
		MachineID:  req.MachineID,
		GVA_MODEL: global.GVA_MODEL{
			CreatedAt: createdAt,
			UpdatedAt: time.Now(),
		},
		CreatedBy: uint(*req.MachineID),
	}, nil
}

func (dataApi *DataApi) CreateData(c *gin.Context) {
	var req CreateDataReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	data, err := req.ToData()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := dataService.CreateData(&data); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
		myMachineApi := MyMachineApi{}
		myMachineApi.UploadDataHook(data)
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
