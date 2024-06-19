// 自动生成模板MachineWarning
package Customize

import (
	"my-server/global"
)

// 机器告警 结构体  MachineWarning
type MachineWarning struct {
	global.GVA_MODEL
	Description string   `json:"description" form:"description" gorm:"column:description;comment:;" binding:"required"`   //描述
	ReporterID  string   `json:"reporterID" form:"reporterID" gorm:"column:reporter_i_d;comment:;" binding:"required"`    //告警联系人ID
	DataTypeID  *int     `json:"dataTypeID" form:"dataTypeID" gorm:"column:data_type_i_d;comment:;" binding:"required"`   //告警数据类型
	Limit       *float64 `json:"limit" form:"limit" gorm:"column:limit;comment:;" binding:"required"`                     //告警阈值
	MachineID   *int     `json:"machineID" form:"machineID" gorm:"column:machine_i_d;comment:;" binding:"required"`       //告警机器ID
	Type        *int     `json:"type" form:"type" gorm:"column:type;comment:0表示大于该阈值时报警，1表示小于该阈值时报警;" binding:"required"` //比较类型
	CreatedBy   uint     `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint     `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint     `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 机器告警 MachineWarning自定义表名 machine_warning
func (MachineWarning) TableName() string {
	return "machine_warning"
}
