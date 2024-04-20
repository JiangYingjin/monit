package Customize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Customize"
	CustomizeReq "github.com/flipped-aurora/gin-vue-admin/server/model/Customize/request"
	"gorm.io/gorm"
)

type MachineService struct {
}

// CreateMachine 创建Machine记录
// Author [piexlmax](https://github.com/piexlmax)
func (machineService *MachineService) CreateMachine(machine *Customize.Machine) (err error) {
	err = global.GVA_DB.Create(machine).Error
	return err
}

// DeleteMachine 删除Machine记录
// Author [piexlmax](https://github.com/piexlmax)
func (machineService *MachineService) DeleteMachine(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Customize.Machine{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Unscoped().Delete(&Customize.Machine{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteMachineByIds 批量删除Machine记录
// Author [piexlmax](https://github.com/piexlmax)
func (machineService *MachineService) DeleteMachineByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Customize.Machine{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&Customize.Machine{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateMachine 更新Machine记录
// Author [piexlmax](https://github.com/piexlmax)
func (machineService *MachineService) UpdateMachine(machine Customize.Machine) (err error) {
	err = global.GVA_DB.Save(&machine).Error
	return err
}

// GetMachine 根据ID获取Machine记录
// Author [piexlmax](https://github.com/piexlmax)
func (machineService *MachineService) GetMachine(ID string) (machine Customize.Machine, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&machine).Error
	return
}

// GetMachineInfoList 分页获取Machine记录
// Author [piexlmax](https://github.com/piexlmax)
func (machineService *MachineService) GetMachineInfoList(info CustomizeReq.MachineSearch) (list []Customize.Machine, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Customize.Machine{})
	var machines []Customize.Machine
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

	err = db.Find(&machines).Error
	return machines, total, err
}
