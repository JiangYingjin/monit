package Customize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Customize"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type MyMachineApi struct {
}

func (m *MyMachineApi) Login(c *gin.Context) {
	response.Ok(c)
}

type GetDataReq struct {
	DataTypeID string   `json:"data_type_ids" form:"data_type_ids" binding:"required"`
	MachineIDs []string `json:"machine_ids" form:"machine_ids" binding:"required"`
	StartTime  string   `json:"start_time" form:"start_time" binding:"required"`
	EndTime    string   `json:"end_time" form:"end_time" binding:"required"`
}

type GetDataRsp struct {
	Data map[string][]Customize.Data `json:"data"` // key: machine_id
}

func (m *MyMachineApi) GetData(c *gin.Context) {
	var req GetDataReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	for _, machineID := range req.MachineIDs {
		_, err := machineService.GetMachine(machineID)
		if err != nil {
			response.FailWithMessage("machine not found", c)
			return
		}
	}

	result := make(map[string][]Customize.Data)
	for _, machineID := range req.MachineIDs {
		tmp := make([]Customize.Data, 0)
		global.GVA_DB.Where("machine_id in ?", req.MachineIDs).Where("data_type_id in ?", req.DataTypeID).Find(&Customize.Data{})
		global.GVA_DB.Where("machine_id = ?", machineID).Where("data_type_id = ?", req.DataTypeID).Where("created_at between ? and ?", req.StartTime, req.EndTime).Find(&tmp)
		result[machineID] = tmp
	}

	response.OkWithData(GetDataRsp{Data: result}, c)
}
