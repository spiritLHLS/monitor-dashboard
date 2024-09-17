package findallpd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/findallpd"
	findallpdReq "github.com/flipped-aurora/gin-vue-admin/server/model/findallpd/request"
	"gorm.io/gorm"
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
// 参数: info findallpdReq.FindallpdSearch - 包含搜索条件和分页信息的结构体
// 返回: list []findallpd.Findallpd - 查找所有产品的列表
//
//	total int64 - 总记录数
//	err error - 错误信息
func (fapdService *FindallpdService) GetFindallpdInfoList(info findallpdReq.FindallpdSearch) (list []findallpd.Findallpd, total int64, err error) {
	// 计算分页参数
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建数据库查询对象
	db := global.GVA_DB.Model(&findallpd.Findallpd{})
	// 构建查询条件
	db = fapdService.buildQueryConditions(db, info)
	// 获取总记录数
	if err = db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	// 应用分页
	if limit > 0 {
		db = db.Limit(limit).Offset(offset)
	}
	// 执行查询
	var findallpds []findallpd.Findallpd
	err = db.Find(&findallpds).Error
	return findallpds, total, err
}

// buildQueryConditions 构建查询条件
// 参数: db *gorm.DB - 数据库查询对象
//
//	info findallpdReq.FindallpdSearch - 搜索条件
//
// 返回: *gorm.DB - 添加了查询条件的数据库查询对象
func (fapdService *FindallpdService) buildQueryConditions(db *gorm.DB, info findallpdReq.FindallpdSearch) *gorm.DB {
	// 时间范围查询
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	// 字符串字段模糊查询
	likeFields := map[string]string{
		"tag": info.Tag, "flag": info.Flag, "billing_type": info.BillingType,
		"memory": info.Memory, "disk": info.Disk, "traffic": info.Traffic,
		"port_speed": info.PortSpeed, "location": info.Location, "price": info.Price,
		"additional": info.Additional, "url": info.Url,
	}
	for field, value := range likeFields {
		if value != "" {
			db = db.Where(field+" LIKE ?", "%"+value+"%")
		}
	}
	// 精确匹配查询
	equalFields := map[string]interface{}{
		"pd_id": info.PdId, "end_id": info.EndId, "cpu": info.Cpu,
		"message_id": info.MessageId,
	}
	for field, value := range equalFields {
		if value != nil && value != "" {
			db = db.Where(field+" = ?", value)
		}
	}
	// 大于查询
	if info.OldStock != nil {
		db = db.Where("old_stock > ?", info.OldStock)
	}
	if info.Stock != nil {
		db = db.Where("stock > ?", info.Stock)
	}
	return db
}

func (fapdService *FindallpdService) GetFindallpdPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
