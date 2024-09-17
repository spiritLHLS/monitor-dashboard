package shops

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shops"
	shopsReq "github.com/flipped-aurora/gin-vue-admin/server/model/shops/request"
	"gorm.io/gorm"
)

type ShopsService struct{}

// CreateShops 创建shops表记录
// Author [piexlmax](https://github.com/piexlmax)
func (spService *ShopsService) CreateShops(sp *shops.Shops) (err error) {
	err = global.GVA_DB.Create(sp).Error
	return err
}

// DeleteShops 删除shops表记录
// Author [piexlmax](https://github.com/piexlmax)
func (spService *ShopsService) DeleteShops(ID string) (err error) {
	err = global.GVA_DB.Delete(&shops.Shops{}, "id = ?", ID).Error
	return err
}

// DeleteShopsByIds 批量删除shops表记录
// Author [piexlmax](https://github.com/piexlmax)
func (spService *ShopsService) DeleteShopsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]shops.Shops{}, "id in ?", IDs).Error
	return err
}

// UpdateShops 更新shops表记录
// Author [piexlmax](https://github.com/piexlmax)
func (spService *ShopsService) UpdateShops(sp shops.Shops) (err error) {
	err = global.GVA_DB.Model(&shops.Shops{}).Where("id = ?", sp.ID).Updates(&sp).Error
	return err
}

// GetShops 根据ID获取shops表记录
// Author [piexlmax](https://github.com/piexlmax)
func (spService *ShopsService) GetShops(ID string) (sp shops.Shops, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&sp).Error
	return
}

// GetShopsInfoList 分页获取shops表记录
// 参数: info shopsReq.ShopsSearch - 包含搜索条件和分页信息的结构体
// 返回: list []shops.Shops - 商店列表
//
//	total int64 - 总记录数
//	err error - 错误信息
func (spService *ShopsService) GetShopsInfoList(info shopsReq.ShopsSearch) (list []shops.Shops, total int64, err error) {
	// 计算分页参数
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建数据库查询对象
	db := global.GVA_DB.Model(&shops.Shops{})
	// 构建查询条件
	db = spService.buildQueryConditions(db, info)
	// 获取总记录数
	if err = db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	// 应用分页
	if limit > 0 {
		db = db.Limit(limit).Offset(offset)
	}
	// 执行查询
	var shopsList []shops.Shops
	err = db.Find(&shopsList).Error
	return shopsList, total, err
}

// buildQueryConditions 构建查询条件
// 参数: db *gorm.DB - 数据库查询对象
//
//	info shopsReq.ShopsSearch - 搜索条件
//
// 返回: *gorm.DB - 添加了查询条件的数据库查询对象
func (spService *ShopsService) buildQueryConditions(db *gorm.DB, info shopsReq.ShopsSearch) *gorm.DB {
	// 时间范围查询
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	// 字符串字段模糊查询
	likeFields := map[string]string{
		"tag":             info.Tag,
		"type":            info.Type,
		"aff_link":        info.AffLink,
		"shop_link":       info.ShopLink,
		"additional_tags": info.AdditionalTags,
	}
	for field, value := range likeFields {
		if value != "" {
			db = db.Where(field+" LIKE ?", "%"+value+"%")
		}
	}
	// 精确匹配查询
	if info.Mid != "" {
		db = db.Where("mid = ?", info.Mid)
	}
	return db
}

func (spService *ShopsService) GetShopsPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
