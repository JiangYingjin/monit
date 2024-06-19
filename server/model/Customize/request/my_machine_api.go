package request

import (
	"my-server/model/Customize"
)

type GetDataReq struct {
	DataTypeID string   `json:"data_type_id" form:"data_type_id" binding:"required"`
	MachineIDs []string `json:"machine_ids" form:"machine_ids" binding:"required"`
	StartTime  string   `json:"start_time" form:"start_time" binding:"required"`
	EndTime    string   `json:"end_time" form:"end_time" binding:"required"`
}

type GetDataRsp struct {
	Data map[string][]Customize.Data `json:"data"` // key: machine_id
}

type MachineLoginReq struct {
	MachineID string `json:"machineID"` // 机器ID
	Password  string `json:"password"`  // 密码
}
