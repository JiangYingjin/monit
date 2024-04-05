// 自动生成模板Data
package Customize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Data 结构体  Data
type Data struct {
	global.GVA_MODEL
	Id         *int   `json:"id" form:"id" gorm:"primarykey;column:id;comment:;" binding:"required"`                 //id
	DataTypeID *int   `json:"name" form:"name" gorm:"column:name;comment:;" binding:"required"`                      //name
	Value      string `json:"description" form:"description" gorm:"column:description;comment:;" binding:"required"` //description
	MachineID  *int   `json:"valueType" form:"valueType" gorm:"column:value_type;comment:;" binding:"required"`      //ValueType
	CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName Data Data自定义表名 data
func (Data) TableName() string {
	return "data"
}
