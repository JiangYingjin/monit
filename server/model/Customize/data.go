// 自动生成模板Data
package Customize

import (
	"my-server/global"
)

// Data 结构体  Data
type Data struct {
	global.GVA_MODEL
	DataTypeID *int     `json:"dataTypeID" form:"dataTypeID" gorm:"column:data_type_i_d;comment:;" binding:"required"` //数据类型
	Value      *float64 `json:"value" form:"value" gorm:"column:value;comment:;" binding:"required"`                   //值
	MachineID  *int     `json:"machineID" form:"machineID" gorm:"column:machine_i_d;comment:;" binding:"required"`     //机器ID
	CreatedBy  uint     `gorm:"column:created_by;comment:创建者"`
	UpdatedBy  uint     `gorm:"column:updated_by;comment:更新者"`
	DeletedBy  uint     `gorm:"column:deleted_by;comment:删除者"`
}

// TableName Data Data自定义表名 data
func (Data) TableName() string {
	return "data"
}
