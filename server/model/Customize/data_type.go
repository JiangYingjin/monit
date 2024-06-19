// 自动生成模板DataType
package Customize

import (
	"my-server/global"
)

// DataType 结构体  DataType
type DataType struct {
	global.GVA_MODEL
	Name        string `json:"name" form:"name" gorm:"column:name;comment:;" binding:"required"`                      //名称
	Description string `json:"description" form:"description" gorm:"column:description;comment:;" binding:"required"` //描述
	Units       string `json:"units" form:"units" gorm:"column:units;comment:;" binding:"required"`                   //单位
	CreatedBy   uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName DataType DataType自定义表名 data_type
func (DataType) TableName() string {
	return "data_type"
}
