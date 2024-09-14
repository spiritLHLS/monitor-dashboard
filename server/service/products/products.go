package products

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/products"
    productsReq "github.com/flipped-aurora/gin-vue-admin/server/model/products/request"
    "gorm.io/gorm"
)

type ProductsService struct {}
// CreateProducts 创建products表记录
// Author [piexlmax](https://github.com/piexlmax)
func (pdService *ProductsService) CreateProducts(pd *products.Products) (err error) {
	err = global.GVA_DB.Create(pd).Error
	return err
}

// DeleteProducts 删除products表记录
// Author [piexlmax](https://github.com/piexlmax)
func (pdService *ProductsService)DeleteProducts(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&products.Products{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&products.Products{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteProductsByIds 批量删除products表记录
// Author [piexlmax](https://github.com/piexlmax)
func (pdService *ProductsService)DeleteProductsByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&products.Products{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&products.Products{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateProducts 更新products表记录
// Author [piexlmax](https://github.com/piexlmax)
func (pdService *ProductsService)UpdateProducts(pd products.Products) (err error) {
	err = global.GVA_DB.Model(&products.Products{}).Where("id = ?",pd.ID).Updates(&pd).Error
	return err
}

// GetProducts 根据ID获取products表记录
// Author [piexlmax](https://github.com/piexlmax)
func (pdService *ProductsService)GetProducts(ID string) (pd products.Products, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&pd).Error
	return
}

// GetProductsInfoList 分页获取products表记录
// Author [piexlmax](https://github.com/piexlmax)
func (pdService *ProductsService)GetProductsInfoList(info productsReq.ProductsSearch) (list []products.Products, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&products.Products{})
    var pds []products.Products
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&pds).Error
	return  pds, total, err
}
func (pdService *ProductsService)GetProductsPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
