package Customize

import (
	"gorm.io/gorm"
	"my-server/global"
	"my-server/model/Customize"
	CustomizeReq "my-server/model/Customize/request"
)

type ServiceTemplateService struct {
}

// CreateServiceTemplate 创建命令模板记录
func (serviceTemplateService *ServiceTemplateService) CreateServiceTemplate(serviceTemplate *Customize.ServiceTemplate) (err error) {
	err = global.GVA_DB.Create(serviceTemplate).Error
	return err
}

// DeleteServiceTemplate 删除命令模板记录
func (serviceTemplateService *ServiceTemplateService) DeleteServiceTemplate(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Customize.ServiceTemplate{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&Customize.ServiceTemplate{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteServiceTemplateByIds 批量删除命令模板记录
func (serviceTemplateService *ServiceTemplateService) DeleteServiceTemplateByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Customize.ServiceTemplate{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&Customize.ServiceTemplate{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateServiceTemplate 更新命令模板记录
func (serviceTemplateService *ServiceTemplateService) UpdateServiceTemplate(serviceTemplate Customize.ServiceTemplate) (err error) {
	err = global.GVA_DB.Save(&serviceTemplate).Error
	return err
}

// GetServiceTemplate 根据ID获取命令模板记录
func (serviceTemplateService *ServiceTemplateService) GetServiceTemplate(ID string) (serviceTemplate Customize.ServiceTemplate, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&serviceTemplate).Error
	return
}

// GetServiceTemplateInfoList 分页获取命令模板记录
func (serviceTemplateService *ServiceTemplateService) GetServiceTemplateInfoList(info CustomizeReq.ServiceTemplateSearch) (list []Customize.ServiceTemplate, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Customize.ServiceTemplate{})
	var serviceTemplates []Customize.ServiceTemplate
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

	err = db.Find(&serviceTemplates).Error
	return serviceTemplates, total, err
}
