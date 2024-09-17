package partitions

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/partitions"
	partitionsReq "github.com/flipped-aurora/gin-vue-admin/server/model/partitions/request"
	"gorm.io/gorm"
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
// 参数: info partitionsReq.PartitionspageSearch - 包含搜索条件和分页信息的结构体
// 返回: list []partitions.Partitionspage - 分区页面列表
//
//	total int64 - 总记录数
//	err error - 错误信息
func (pspService *PartitionspageService) GetPartitionspageInfoList(info partitionsReq.PartitionspageSearch) (list []partitions.Partitionspage, total int64, err error) {
	// 计算分页参数
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建数据库查询对象
	db := global.GVA_DB.Model(&partitions.Partitionspage{})
	// 构建查询条件
	db = pspService.buildQueryConditions(db, info)
	// 获取总记录数
	if err = db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	// 应用分页
	if limit > 0 {
		db = db.Limit(limit).Offset(offset)
	}
	// 执行查询
	var partitionspages []partitions.Partitionspage
	err = db.Find(&partitionspages).Error
	return partitionspages, total, err
}

// buildQueryConditions 构建查询条件
// 参数: db *gorm.DB - 数据库查询对象
//
//	info partitionsReq.PartitionspageSearch - 搜索条件
//
// 返回: *gorm.DB - 添加了查询条件的数据库查询对象
func (pspService *PartitionspageService) buildQueryConditions(db *gorm.DB, info partitionsReq.PartitionspageSearch) *gorm.DB {
	// 时间范围查询
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	// 字符串字段模糊查询
	stringFields := map[string]string{
		"tg_tag":     info.TgTag,
		"name":       info.Name,
		"link":       info.Link,
		"additional": info.Additional,
	}
	for field, value := range stringFields {
		if value != "" {
			db = db.Where(field+" LIKE ?", "%"+value+"%")
		}
	}
	// 精确匹配查询
	if info.Type != "" {
		db = db.Where("type = ?", info.Type)
	}
	if info.Num != nil {
		db = db.Where("num = ?", info.Num)
	}
	// 区间查询
	if info.Intervals != nil {
		db = db.Where("intervals < ?", info.Intervals)
	}
	return db
}

func (pspService *PartitionspageService) GetPartitionspagePublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
