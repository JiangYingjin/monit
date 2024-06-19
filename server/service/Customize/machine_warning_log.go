package Customize

import (
	"my-server/global"
	"my-server/model/Customize"
	CustomizeReq "my-server/model/Customize/request"
)

type MachineWarningLogService struct {
}

// CreateMachineWarningLog 创建machineWarningLog表记录
func (machineWarningLogService *MachineWarningLogService) CreateMachineWarningLog(machineWarningLog *Customize.MachineWarningLog) (err error) {
	err = global.GVA_DB.Create(machineWarningLog).Error
	return err
}

// DeleteMachineWarningLog 删除machineWarningLog表记录
func (machineWarningLogService *MachineWarningLogService) DeleteMachineWarningLog(userId string) (err error) {
	err = global.GVA_DB.Delete(&Customize.MachineWarningLog{}, "user_id = ?", userId).Error
	return err
}

// DeleteMachineWarningLogByIds 批量删除machineWarningLog表记录
func (machineWarningLogService *MachineWarningLogService) DeleteMachineWarningLogByIds(userIds []string) (err error) {
	err = global.GVA_DB.Delete(&[]Customize.MachineWarningLog{}, "user_id in ?", userIds).Error
	return err
}

// UpdateMachineWarningLog 更新machineWarningLog表记录
func (machineWarningLogService *MachineWarningLogService) UpdateMachineWarningLog(machineWarningLog Customize.MachineWarningLog) (err error) {
	err = global.GVA_DB.Save(&machineWarningLog).Error
	return err
}

// GetMachineWarningLog 根据userId获取machineWarningLog表记录
func (machineWarningLogService *MachineWarningLogService) GetMachineWarningLog(userId string) (machineWarningLog Customize.MachineWarningLog, err error) {
	err = global.GVA_DB.Where("user_id = ?", userId).First(&machineWarningLog).Error
	return
}

// GetMachineWarningLogInfoList 分页获取machineWarningLog表记录
func (machineWarningLogService *MachineWarningLogService) GetMachineWarningLogInfoList(info CustomizeReq.MachineWarningLogSearch) (list []Customize.MachineWarningLog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Customize.MachineWarningLog{})
	var machineWarningLogs []Customize.MachineWarningLog
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&machineWarningLogs).Error
	return machineWarningLogs, total, err
}
