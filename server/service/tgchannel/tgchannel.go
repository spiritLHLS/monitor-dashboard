package tgchannel

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tgchannel"
	tgchannelReq "github.com/flipped-aurora/gin-vue-admin/server/model/tgchannel/request"
)

type TgchannelService struct{}

// CreateTgchannel 创建tgchannel表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tgcService *TgchannelService) CreateTgchannel(tgc *tgchannel.Tgchannel) (err error) {
	err = global.GVA_DB.Create(tgc).Error
	return err
}

// DeleteTgchannel 删除tgchannel表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tgcService *TgchannelService) DeleteTgchannel(ID string) (err error) {
	err = global.GVA_DB.Delete(&tgchannel.Tgchannel{}, "id = ?", ID).Error
	return err
}

// DeleteTgchannelByIds 批量删除tgchannel表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tgcService *TgchannelService) DeleteTgchannelByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]tgchannel.Tgchannel{}, "id in ?", IDs).Error
	return err
}

// UpdateTgchannel 更新tgchannel表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tgcService *TgchannelService) UpdateTgchannel(tgc tgchannel.Tgchannel) (err error) {
	err = global.GVA_DB.Model(&tgchannel.Tgchannel{}).Where("id = ?", tgc.ID).Updates(&tgc).Error
	return err
}

// GetTgchannel 根据ID获取tgchannel表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tgcService *TgchannelService) GetTgchannel(ID string) (tgc tgchannel.Tgchannel, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&tgc).Error
	return
}

// GetTgchannelInfoList 分页获取tgchannel表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tgcService *TgchannelService) GetTgchannelInfoList(info tgchannelReq.TgchannelSearch) (list []tgchannel.Tgchannel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&tgchannel.Tgchannel{})
	var tgcs []tgchannel.Tgchannel
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Channel != "" {
		db = db.Where("channel LIKE ?", "%"+info.Channel+"%")
	}
	if info.Tgid != "" {
		db = db.Where("tgid LIKE ?", "%"+info.Tgid+"%")
	}
	if info.Flag != "" {
		db = db.Where("flag LIKE ?", "%"+info.Flag+"%")
	}
	if info.AdditionalKey != "" {
		db = db.Where("additional_key LIKE ?", "%"+info.AdditionalKey+"%")
	}
	if info.Tokens != "" {
		db = db.Where("tokens LIKE ?", "%"+info.Tokens+"%")
	}
	if info.Tags != "" {
		db = db.Where("tags LIKE ?", "%"+info.Tags+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&tgcs).Error
	return tgcs, total, err
}
func (tgcService *TgchannelService) GetTgchannelPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
