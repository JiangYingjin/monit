// 自动生成模板MachineWarning
package Customize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 机器告警 结构体  MachineWarning
type MachineWarning struct {
	global.GVA_MODEL
	Description string   `json:"description" form:"description" gorm:"column:description;comment:;" binding:"required"` //描述
	Limit       *float64 `json:"limit" form:"limit" gorm:"column:limit;comment:;" binding:"required"`                   //告警阈值
	Reporter_id string   `json:"reporter_id" form:"reporter_id" gorm:"column:reporter_id;comment:;" binding:"required"` //告警联系人ID
	CreatedBy   uint     `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint     `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint     `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 机器告警 MachineWarning自定义表名 machine_warning
func (MachineWarning) TableName() string {
	return "machine_warning"
}
