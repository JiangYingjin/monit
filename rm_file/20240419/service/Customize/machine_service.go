package Customize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Customize"
	CustomizeReq "github.com/flipped-aurora/gin-vue-admin/server/model/Customize/request"
	"gorm.io/gorm"
)

type MachineServiceService struct {
}

// CreateMachineService 创建数据类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (machineServiceService *MachineServiceService) CreateMachineService(machineService *Customize.MachineService) (err error) {
	err = global.GVA_DB.Create(machineService).Error
	return err
}

// DeleteMachineService 删除数据类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (machineServiceService *MachineServiceService) DeleteMachineService(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Customize.MachineService{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&Customize.MachineService{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteMachineServiceByIds 批量删除数据类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (machineServiceService *MachineServiceService) DeleteMachineServiceByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Customize.MachineService{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&Customize.MachineService{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateMachineService 更新数据类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (machineServiceService *MachineServiceService) UpdateMachineService(machineService Customize.MachineService) (err error) {
	err = global.GVA_DB.Save(&machineService).Error
	return err
}

// GetMachineService 根据ID获取数据类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (machineServiceService *MachineServiceService) GetMachineService(ID string) (machineService Customize.MachineService, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&machineService).Error
	return
}

// GetMachineServiceInfoList 分页获取数据类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (machineServiceService *MachineServiceService) GetMachineServiceInfoList(info CustomizeReq.MachineServiceSearch) (list []Customize.MachineService, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Customize.MachineService{})
	var machineServices []Customize.MachineService
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

	err = db.Find(&machineServices).Error
	return machineServices, total, err
}
