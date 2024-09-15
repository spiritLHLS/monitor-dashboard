package partitions

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/partitions"
	partitionsReq "github.com/flipped-aurora/gin-vue-admin/server/model/partitions/request"
)

type PartitionspageService struct{}

// CreatePartitionspage 创建partitionspage表记录
// Author [piexlmax](https://github.com/piexlmax)
func (pspService *PartitionspageService) CreatePartitionspage(psp *partitions.Partitionspage) (err error) {
	err = global.GVA_DB.Create(psp).Error
	return err
}

// DeletePartitionspage 删除partitionspage表记录
// Author [piexlmax](https://github.com/piexlmax)
func (pspService *PartitionspageService) DeletePartitionspage(ID string) (err error) {
	err = global.GVA_DB.Delete(&partitions.Partitionspage{}, "id = ?", ID).Error
	return err
}

// DeletePartitionspageByIds 批量删除partitionspage表记录
// Author [piexlmax](https://github.com/piexlmax)
func (pspService *PartitionspageService) DeletePartitionspageByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]partitions.Partitionspage{}, "id in ?", IDs).Error
	return err
}

// UpdatePartitionspage 更新partitionspage表记录
// Author [piexlmax](https://github.com/piexlmax)
func (pspService *PartitionspageService) UpdatePartitionspage(psp partitions.Partitionspage) (err error) {
	err = global.GVA_DB.Model(&partitions.Partitionspage{}).Where("id = ?", psp.ID).Updates(&psp).Error
	return err
}

// GetPartitionspage 根据ID获取partitionspage表记录
// Author [piexlmax](https://github.com/piexlmax)
func (pspService *PartitionspageService) GetPartitionspage(ID string) (psp partitions.Partitionspage, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&psp).Error
	return
}

// GetPartitionspageInfoList 分页获取partitionspage表记录
// Author [piexlmax](https://github.com/piexlmax)
func (pspService *PartitionspageService) GetPartitionspageInfoList(info partitionsReq.PartitionspageSearch) (list []partitions.Partitionspage, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&partitions.Partitionspage{})
	var psps []partitions.Partitionspage
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.TgTag != "" {
		db = db.Where("tg_tag LIKE ?", "%"+info.TgTag+"%")
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Link != "" {
		db = db.Where("link LIKE ?", "%"+info.Link+"%")
	}
	if info.Type != "" {
		db = db.Where("type = ?", info.Type)
	}
	if info.Num != nil {
		db = db.Where("num = ?", info.Num)
	}
	if info.Additional != "" {
		db = db.Where("additional LIKE ?", "%"+info.Additional+"%")
	}
	if info.Intervals != nil {
		db = db.Where("intervals < ?", info.Intervals)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&psps).Error
	return psps, total, err
}
func (pspService *PartitionspageService) GetPartitionspagePublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
