package tgchannel

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tgchannel"
	tgchannelReq "github.com/flipped-aurora/gin-vue-admin/server/model/tgchannel/request"
	"gorm.io/gorm"
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
// 参数: info tgchannelReq.TgchannelSearch - 包含搜索条件和分页信息的结构体
// 返回: list []tgchannel.Tgchannel - Telegram频道列表
//
//	total int64 - 总记录数
//	err error - 错误信息
func (tgcService *TgchannelService) GetTgchannelInfoList(info tgchannelReq.TgchannelSearch) (list []tgchannel.Tgchannel, total int64, err error) {
	// 计算分页参数
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建数据库查询对象
	db := global.GVA_DB.Model(&tgchannel.Tgchannel{})
	// 构建查询条件
	db = tgcService.buildQueryConditions(db, info)
	// 获取总记录数
	if err = db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	// 应用分页
	if limit > 0 {
		db = db.Limit(limit).Offset(offset)
	}
	// 执行查询
	var tgChannels []tgchannel.Tgchannel
	err = db.Find(&tgChannels).Error
	return tgChannels, total, err
}

// buildQueryConditions 构建查询条件
// 参数: db *gorm.DB - 数据库查询对象
//
//	info tgchannelReq.TgchannelSearch - 搜索条件
//
// 返回: *gorm.DB - 添加了查询条件的数据库查询对象
func (tgcService *TgchannelService) buildQueryConditions(db *gorm.DB, info tgchannelReq.TgchannelSearch) *gorm.DB {
	// 时间范围查询
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	// 字符串字段模糊查询
	likeFields := map[string]string{
		"channel":        info.Channel,
		"tgid":           info.Tgid,
		"flag":           info.Flag,
		"additional_key": info.AdditionalKey,
		"tokens":         info.Tokens,
		"tags":           info.Tags,
	}
	for field, value := range likeFields {
		if value != "" {
			db = db.Where(field+" LIKE ?", "%"+value+"%")
		}
	}
	return db
}

func (tgcService *TgchannelService) GetTgchannelPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
