// 自动生成模板MachineService
package Customize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 数据类型 结构体  MachineService
type MachineService struct {
	global.GVA_MODEL
	MachineID *int   `json:"machineID" form:"machineID" gorm:"column:machine_i_d;comment:;" binding:"required"`      //机器ID
	Services  string `json:"services" form:"services" gorm:"column:services;comment:;type:text;" binding:"required"` //当前机器支持的服务
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 数据类型 MachineService自定义表名 machine_service
func (MachineService) TableName() string {
	return "machine_service"
}
