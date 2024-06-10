// 自动生成模板MachineWarningLog
package Customize

import (
	"time"
)

// machineWarningLog表 结构体  MachineWarningLog
type MachineWarningLog struct {
	UserId    *int       `json:"userId" form:"userId" gorm:"primarykey;column:user_id;comment:;size:20;"`          //userId字段
	WarningId *int       `json:"warningId" form:"warningId" gorm:"primarykey;column:warning_id;comment:;size:20;"` //warningId字段
	SendTime  *time.Time `json:"sendTime" form:"sendTime" gorm:"primarykey;column:send_time;comment:;"`            //sendTime字段
}

// TableName machineWarningLog表 MachineWarningLog自定义表名 machine_warning_log
func (MachineWarningLog) TableName() string {
	return "machine_warning_log"
}
