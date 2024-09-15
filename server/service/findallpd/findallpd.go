package findallpd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/findallpd"
	findallpdReq "github.com/flipped-aurora/gin-vue-admin/server/model/findallpd/request"
)

type FindallpdService struct{}

// CreateFindallpd 创建findallpd表记录
// Author [piexlmax](https://github.com/piexlmax)
func (fapdService *FindallpdService) CreateFindallpd(fapd *findallpd.Findallpd) (err error) {
	err = global.GVA_DB.Create(fapd).Error
	return err
}

// DeleteFindallpd 删除findallpd表记录
// Author [piexlmax](https://github.com/piexlmax)
func (fapdService *FindallpdService) DeleteFindallpd(ID string) (err error) {
	err = global.GVA_DB.Delete(&findallpd.Findallpd{}, "id = ?", ID).Error
	return err
}

// DeleteFindallpdByIds 批量删除findallpd表记录
// Author [piexlmax](https://github.com/piexlmax)
func (fapdService *FindallpdService) DeleteFindallpdByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]findallpd.Findallpd{}, "id in ?", IDs).Error
	return err
}

// UpdateFindallpd 更新findallpd表记录
// Author [piexlmax](https://github.com/piexlmax)
func (fapdService *FindallpdService) UpdateFindallpd(fapd findallpd.Findallpd) (err error) {
	err = global.GVA_DB.Model(&findallpd.Findallpd{}).Where("id = ?", fapd.ID).Updates(&fapd).Error
	return err
}

// GetFindallpd 根据ID获取findallpd表记录
// Author [piexlmax](https://github.com/piexlmax)
func (fapdService *FindallpdService) GetFindallpd(ID string) (fapd findallpd.Findallpd, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&fapd).Error
	return
}

// GetFindallpdInfoList 分页获取findallpd表记录
// Author [piexlmax](https://github.com/piexlmax)
func (fapdService *FindallpdService) GetFindallpdInfoList(info findallpdReq.FindallpdSearch) (list []findallpd.Findallpd, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&findallpd.Findallpd{})
	var fapds []findallpd.Findallpd
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Tag != "" {
		db = db.Where("tag LIKE ?", "%"+info.Tag+"%")
	}
	if info.Flag != "" {
		db = db.Where("flag LIKE ?", "%"+info.Flag+"%")
	}
	if info.BillingType != "" {
		db = db.Where("billing_type LIKE ?", "%"+info.BillingType+"%")
	}
	if info.PdId != nil {
		db = db.Where("pd_id = ?", info.PdId)
	}
	if info.EndId != nil {
		db = db.Where("end_id = ?", info.EndId)
	}
	if info.Cpu != "" {
		db = db.Where("cpu = ?", info.Cpu)
	}
	if info.Memory != "" {
		db = db.Where("memory LIKE ?", "%"+info.Memory+"%")
	}
	if info.Disk != "" {
		db = db.Where("disk LIKE ?", "%"+info.Disk+"%")
	}
	if info.Traffic != "" {
		db = db.Where("traffic LIKE ?", "%"+info.Traffic+"%")
	}
	if info.PortSpeed != "" {
		db = db.Where("port_speed LIKE ?", "%"+info.PortSpeed+"%")
	}
	if info.Location != "" {
		db = db.Where("location LIKE ?", "%"+info.Location+"%")
	}
	if info.Price != "" {
		db = db.Where("price LIKE ?", "%"+info.Price+"%")
	}
	if info.Additional != "" {
		db = db.Where("additional LIKE ?", "%"+info.Additional+"%")
	}
	if info.Url != "" {
		db = db.Where("url LIKE ?", "%"+info.Url+"%")
	}
	if info.OldStock != nil {
		db = db.Where("old_stock > ?", info.OldStock)
	}
	if info.Stock != nil {
		db = db.Where("stock > ?", info.Stock)
	}
	if info.MessageId != "" {
		db = db.Where("message_id = ?", info.MessageId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&fapds).Error
	return fapds, total, err
}
func (fapdService *FindallpdService) GetFindallpdPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
