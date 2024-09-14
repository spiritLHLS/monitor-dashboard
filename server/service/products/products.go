package products

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/products"
	productsReq "github.com/flipped-aurora/gin-vue-admin/server/model/products/request"
)

type ProductsService struct{}

// CreateProducts 创建products表记录
// Author [piexlmax](https://github.com/piexlmax)
func (pdService *ProductsService) CreateProducts(pd *products.Products) (err error) {
	err = global.GVA_DB.Create(pd).Error
	return err
}

// DeleteProducts 删除products表记录
// Author [piexlmax](https://github.com/piexlmax)
func (pdService *ProductsService) DeleteProducts(ID string) (err error) {
	err = global.GVA_DB.Delete(&products.Products{}, "id = ?", ID).Error
	return err
}

// DeleteProductsByIds 批量删除products表记录
// Author [piexlmax](https://github.com/piexlmax)
func (pdService *ProductsService) DeleteProductsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]products.Products{}, "id in ?", IDs).Error
	return err
}

// UpdateProducts 更新products表记录
// Author [piexlmax](https://github.com/piexlmax)
func (pdService *ProductsService) UpdateProducts(pd products.Products) (err error) {
	err = global.GVA_DB.Model(&products.Products{}).Where("id = ?", pd.ID).Updates(&pd).Error
	return err
}

// GetProducts 根据ID获取products表记录
// Author [piexlmax](https://github.com/piexlmax)
func (pdService *ProductsService) GetProducts(ID string) (pd products.Products, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&pd).Error
	return
}

// GetProductsInfoList 分页获取商品记录
// Author [piexlmax](https://github.com/piexlmax)
func (pdService *ProductsService) GetProductsInfoList(info productsReq.ProductsSearch) (list []products.Products, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&products.Products{})
	var pds []products.Products
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Tag != "" {
		db = db.Where("tag LIKE ?", "%"+info.Tag+"%")
	}
	if info.Cpu != "" {
		db = db.Where("cpu LIKE ?", "%"+info.Cpu+"%")
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
	if info.BillingType != "" {
		db = db.Where("billing_type LIKE ?", "%"+info.BillingType+"%")
	}
	if info.OldStock != nil {
		db = db.Where("old_stock > ?", info.OldStock)
	}
	if info.Stock != nil {
		db = db.Where("stock > ?", info.Stock)
	}
	if info.MultiCheck != nil {
		db = db.Where("multiCheck = ?", info.Intervals)
	}
	if info.Intervals != nil {
		db = db.Where("intervals < ?", info.Intervals)
	}
	if info.MessageId != "" {
		db = db.Where("message_id = ?", info.MessageId)
	}
	if info.PushIntervals != nil {
		db = db.Where("push_intervals < ?", info.PushIntervals)
	}
	if info.PushTime != nil {
		db = db.Where("push_time > ?", info.PushTime)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&pds).Error
	return pds, total, err
}

func (pdService *ProductsService) GetProductsPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
