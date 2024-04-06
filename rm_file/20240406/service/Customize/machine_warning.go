package Customize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Customize"
	CustomizeReq "github.com/flipped-aurora/gin-vue-admin/server/model/Customize/request"
	"gorm.io/gorm"
)

type MachineWarningService struct {
}

// CreateMachineWarning 创建机器告警记录
// Author [piexlmax](https://github.com/piexlmax)
func (machineWarningService *MachineWarningService) CreateMachineWarning(machineWarning *Customize.MachineWarning) (err error) {
	err = global.GVA_DB.Create(machineWarning).Error
	return err
}

// DeleteMachineWarning 删除机器告警记录
// Author [piexlmax](https://github.com/piexlmax)
func (machineWarningService *MachineWarningService) DeleteMachineWarning(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Customize.MachineWarning{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&Customize.MachineWarning{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteMachineWarningByIds 批量删除机器告警记录
// Author [piexlmax](https://github.com/piexlmax)
func (machineWarningService *MachineWarningService) DeleteMachineWarningByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Customize.MachineWarning{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&Customize.MachineWarning{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateMachineWarning 更新机器告警记录
// Author [piexlmax](https://github.com/piexlmax)
func (machineWarningService *MachineWarningService) UpdateMachineWarning(machineWarning Customize.MachineWarning) (err error) {
	err = global.GVA_DB.Save(&machineWarning).Error
	return err
}

// GetMachineWarning 根据ID获取机器告警记录
// Author [piexlmax](https://github.com/piexlmax)
func (machineWarningService *MachineWarningService) GetMachineWarning(ID string) (machineWarning Customize.MachineWarning, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&machineWarning).Error
	return
}

// GetMachineWarningInfoList 分页获取机器告警记录
// Author [piexlmax](https://github.com/piexlmax)
func (machineWarningService *MachineWarningService) GetMachineWarningInfoList(info CustomizeReq.MachineWarningSearch) (list []Customize.MachineWarning, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Customize.MachineWarning{})
	var machineWarnings []Customize.MachineWarning
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&machineWarnings).Error
	return machineWarnings, total, err
}
