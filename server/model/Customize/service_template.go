// 自动生成模板ServiceTemplate
package Customize

import (
	"my-server/global"
)

// 命令模板 结构体  ServiceTemplate
type ServiceTemplate struct {
	global.GVA_MODEL
	Service   string `json:"service" form:"service" gorm:"column:service;comment:;" binding:"required"`              //服务名
	Template  string `json:"template" form:"template" gorm:"type:json;column:template;comment:;" binding:"required"` //命令模板
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 命令模板 ServiceTemplate自定义表名 service_template
func (ServiceTemplate) TableName() string {
	return "service_template"
}
