// 自动生成模板DataType
package Customize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// DataType 结构体  DataType
type DataType struct {
	global.GVA_MODEL
	Id          *int   `json:"id" form:"id" gorm:"primarykey;column:id;comment:;" binding:"required"`                 //id
	Name        string `json:"name" form:"name" gorm:"column:name;comment:;" binding:"required"`                      //name
	Description string `json:"description" form:"description" gorm:"column:description;comment:;" binding:"required"` //description
	CreatedBy   uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName DataType DataType自定义表名 data_type
func (DataType) TableName() string {
	return "data_type"
}
