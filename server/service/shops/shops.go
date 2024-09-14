package shops

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shops"
	shopsReq "github.com/flipped-aurora/gin-vue-admin/server/model/shops/request"
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
// Author [piexlmax](https://github.com/piexlmax)
func (spService *ShopsService) GetShopsInfoList(info shopsReq.ShopsSearch) (list []shops.Shops, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&shops.Shops{})
	var sps []shops.Shops
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Tag != "" {
		db = db.Where("tag LIKE ?", "%"+info.Tag+"%")
	}
	if info.Mid != "" {
		db = db.Where("mid = ?", info.Mid)
	}
	if info.Type != "" {
		db = db.Where("type LIKE ?", "%"+info.Type+"%")
	}
	if info.AffLink != "" {
		db = db.Where("aff_link LIKE ?", "%"+info.AffLink+"%")
	}
	if info.ShopLink != "" {
		db = db.Where("shop_link LIKE ?", "%"+info.ShopLink+"%")
	}
	if info.AdditionalTags != "" {
		db = db.Where("additional_tags LIKE ?", "%"+info.AdditionalTags+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&sps).Error
	return sps, total, err
}
func (spService *ShopsService) GetShopsPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
