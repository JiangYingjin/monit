// 自动生成模板Machine
package Customize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Machine 结构体  Machine
type Machine struct {
	global.GVA_MODEL
	Name        string `json:"name" form:"name" gorm:"column:name;comment:;" binding:"required"`                      //名字
	Description string `json:"description" form:"description" gorm:"column:description;comment:;" binding:"required"` //描述
	IPAddr      string `json:"ip_addr" form:"ip_addr" gorm:"column:ip_addr;comment:;" binding:"required"`             //IP地址
	Password    string `json:"password" form:"password" gorm:"column:password;comment:;" binding:"required"`          //密钥
	Status      *bool  `json:"status" form:"status" gorm:"column:status;comment:机器是否在线;"`                             //在线状态
	Service     string `json:"service" form:"service" gorm:"type:json;column:service;comment:当前机器支持监听的服务;"`           //服务
	CreatedBy   uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName Machine Machine自定义表名 machine
func (Machine) TableName() string {
	return "machine"
}
