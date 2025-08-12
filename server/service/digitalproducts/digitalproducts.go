package digitalproducts

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/digitalproducts"
	digitalproductsReq "github.com/flipped-aurora/gin-vue-admin/server/model/digitalproducts/request"
	"gorm.io/gorm"
)

type DigitalProductsService struct{}

// CreateDigitalProducts 创建数字商品表记录
// Author [yourname](https://github.com/yourname)
func (dpdService *DigitalProductsService) CreateDigitalProducts(dpd *digitalproducts.DigitalProducts) (err error) {
	err = global.GVA_DB.Create(dpd).Error
	return err
}

// DeleteDigitalProducts 删除数字商品表记录
// Author [yourname](https://github.com/yourname)
func (dpdService *DigitalProductsService) DeleteDigitalProducts(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&digitalproducts.DigitalProducts{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&digitalproducts.DigitalProducts{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteDigitalProductsByIds 批量删除数字商品表记录
// Author [yourname](https://github.com/yourname)
func (dpdService *DigitalProductsService) DeleteDigitalProductsByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&digitalproducts.DigitalProducts{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&digitalproducts.DigitalProducts{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateDigitalProducts 更新数字商品表记录
// Author [yourname](https://github.com/yourname)
func (dpdService *DigitalProductsService) UpdateDigitalProducts(dpd digitalproducts.DigitalProducts) (err error) {
	err = global.GVA_DB.Model(&digitalproducts.DigitalProducts{}).Where("id = ?", dpd.ID).Updates(&dpd).Error
	return err
}

// GetDigitalProducts 根据ID获取数字商品表记录
// Author [yourname](https://github.com/yourname)
func (dpdService *DigitalProductsService) GetDigitalProducts(ID string) (dpd digitalproducts.DigitalProducts, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&dpd).Error
	return
}

// GetDigitalProductsInfoList 分页获取数字商品表记录
// Author [yourname](https://github.com/yourname)
func (dpdService *DigitalProductsService) GetDigitalProductsInfoList(info digitalproductsReq.DigitalProductsSearch) (list []digitalproducts.DigitalProducts, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&digitalproducts.DigitalProducts{})
	var dpds []digitalproducts.DigitalProducts
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.PriceUnit != nil && *info.PriceUnit != "" {
		db = db.Where("price_unit = ?", *info.PriceUnit)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&dpds).Error
	return dpds, total, err
}
func (dpdService *DigitalProductsService) GetDigitalProductsPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
