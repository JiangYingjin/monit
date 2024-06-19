package Customize

import (
	"gorm.io/gorm"
	"my-server/global"
	"my-server/model/Customize"
	CustomizeReq "my-server/model/Customize/request"
)

type DataTypeService struct {
}

// CreateDataType 创建DataType记录
func (dataTypeService *DataTypeService) CreateDataType(dataType *Customize.DataType) (err error) {
	err = global.GVA_DB.Create(dataType).Error
	return err
}

// DeleteDataType 删除DataType记录
func (dataTypeService *DataTypeService) DeleteDataType(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Customize.DataType{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&Customize.DataType{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteDataTypeByIds 批量删除DataType记录
func (dataTypeService *DataTypeService) DeleteDataTypeByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Customize.DataType{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&Customize.DataType{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateDataType 更新DataType记录
func (dataTypeService *DataTypeService) UpdateDataType(dataType Customize.DataType) (err error) {
	err = global.GVA_DB.Save(&dataType).Error
	return err
}

// GetDataType 根据ID获取DataType记录
func (dataTypeService *DataTypeService) GetDataType(ID string) (dataType Customize.DataType, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&dataType).Error
	return
}

// GetDataTypeInfoList 分页获取DataType记录
func (dataTypeService *DataTypeService) GetDataTypeInfoList(info CustomizeReq.DataTypeSearch) (list []Customize.DataType, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Customize.DataType{})
	var dataTypes []Customize.DataType
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

	err = db.Find(&dataTypes).Error
	return dataTypes, total, err
}
