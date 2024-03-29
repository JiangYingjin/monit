// 自动生成模板Machine
package Customize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Machine 结构体  Machine
type Machine struct {
	global.GVA_MODEL
	Id          *int   `json:"id" form:"id" gorm:"primarykey;column:id;comment:;" binding:"required"`                 //id
	Name        string `json:"name" form:"name" gorm:"column:name;comment:;" binding:"required"`                      //name
	Description string `json:"description" form:"description" gorm:"column:description;comment:;" binding:"required"` //description
	IPAddr      string `json:"valueType" form:"valueType" gorm:"column:value_type;comment:;" binding:"required"`      //ValueType
	CreatedBy   uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName Machine Machine自定义表名 machine
func (Machine) TableName() string {
	return "machine"
}
