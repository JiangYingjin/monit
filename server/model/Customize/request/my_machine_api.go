package request

import "time"

type MachineDataReq struct {
	MachineIDs []string   `json:"machineIDs" form:"machineIDs"`
	StartTime  *time.Time `json:"startTime" form:"startTime"`
	EndTime    *time.Time `json:"endTime" form:"endTime"`
}

type MachineLoginReq struct {
	MachineID string `json:"machineID"` // 机器ID
	Password  string `json:"password"`  // 密码
}
