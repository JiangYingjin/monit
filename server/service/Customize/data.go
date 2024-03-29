package Customize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Customize"
	CustomizeReq "github.com/flipped-aurora/gin-vue-admin/server/model/Customize/request"
	"gorm.io/gorm"
)

type DataService struct {
}

// CreateData 创建Data记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataService *DataService) CreateData(data *Customize.Data) (err error) {
	err = global.GVA_DB.Create(data).Error
	return err
}

// DeleteData 删除Data记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataService *DataService) DeleteData(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Customize.Data{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&Customize.Data{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteDataByIds 批量删除Data记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataService *DataService) DeleteDataByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Customize.Data{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&Customize.Data{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateData 更新Data记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataService *DataService) UpdateData(data Customize.Data) (err error) {
	err = global.GVA_DB.Save(&data).Error
	return err
}

// GetData 根据ID获取Data记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataService *DataService) GetData(ID string) (data Customize.Data, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&data).Error
	return
}

// GetDataInfoList 分页获取Data记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataService *DataService) GetDataInfoList(info CustomizeReq.DataSearch) (list []Customize.Data, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Customize.Data{})
	var datas []Customize.Data
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

	err = db.Find(&datas).Error
	return datas, total, err
}
