package products

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/products"
	productsReq "github.com/flipped-aurora/gin-vue-admin/server/model/products/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shops"
	"gorm.io/gorm"
	"strings"
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
// 参数: info productsReq.ProductsSearch - 包含搜索条件和分页信息的结构体
// 返回: list []products.Products - 商品列表
//
//	total int64 - 总记录数
//	err error - 错误信息
func (pdService *ProductsService) GetProductsInfoList(info productsReq.ProductsSearch) (list []products.Products, total int64, err error) {
	// 计算分页参数
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建数据库查询对象
	db := global.GVA_DB.Model(&products.Products{})
	// 构建查询条件
	db = pdService.buildQueryConditions(db, info)
	// 获取总记录数
	if err = db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	// 应用分页
	if limit > 0 {
		db = db.Limit(limit).Offset(offset)
	}
	// 执行查询
	var pdProducts []products.Products
	err = db.Find(&pdProducts).Error
	return pdProducts, total, err
}

// buildQueryConditions 构建查询条件
// 参数: db *gorm.DB - 数据库查询对象
//
//	info productsReq.ProductsSearch - 搜索条件
//
// 返回: *gorm.DB - 添加了查询条件的数据库查询对象
func (pdService *ProductsService) buildQueryConditions(db *gorm.DB, info productsReq.ProductsSearch) *gorm.DB {
	// 时间范围查询
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	// 字符串字段模糊查询
	stringFields := map[string]string{
		"tag": info.Tag, "cpu": info.Cpu, "memory": info.Memory,
		"disk": info.Disk, "traffic": info.Traffic, "port_speed": info.PortSpeed,
		"location": info.Location, "price": info.Price, "additional": info.Additional,
		"url": info.Url, "billing_type": info.BillingType,
	}
	for field, value := range stringFields {
		if value != "" {
			db = db.Where(field+" LIKE ?", "%"+value+"%")
		}
	}
	// 数值字段精确查询
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
	if info.PushIntervals != nil {
		db = db.Where("push_intervals < ?", info.PushIntervals)
	}
	// 其他字段查询
	if info.MessageId != "" {
		db = db.Where("message_id = ?", info.MessageId)
	}
	if info.PushTime != nil {
		db = db.Where("push_time > ?", info.PushTime)
	}
	return db
}

// GetProductsPublic 获取公开的商品信息列表
// 参数: info productsReq.ProductsSearch - 包含搜索条件和分页信息的结构体
// 返回: list []products.Products - 商品列表
//
//	total int64 - 总记录数
//	err error - 错误信息
func (pdService *ProductsService) GetProductsPublic(info productsReq.PublicProductsSearch) (list []products.Products, total int64, err error) {
	// 计算分页参数
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建数据库查询对象
	db := global.GVA_DB.Model(&products.Products{})
	// 构建查询条件
	db = pdService.buildPublicQueryConditions(db, info)
	// 获取总记录数
	if err = db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	// 应用分页
	if limit > 0 {
		db = db.Limit(limit).Offset(offset)
	}
	// 执行查询
	var pdProducts []products.Products
	err = db.Find(&pdProducts).Error
	// 优化：使用map缓存shop信息
	shopCache := make(map[string]*shops.Shops)
	// 重构URL字段带上商家信息
	if total > 0 {
		for i := range pdProducts {
			pd := &pdProducts[i]
			// 检查缓存中是否已有该tag的shop信息
			shop, exists := shopCache[pd.Tag]
			if !exists {
				// 如果缓存中没有，则查询数据库
				shop = &shops.Shops{}
				sdb := global.GVA_DB.Model(&shops.Shops{})
				sdb = sdb.Where("tag = ?", pd.Tag)
				shopErr := sdb.First(shop).Error
				if shopErr != nil {
					continue
				}
				// 将查询结果存入缓存
				shopCache[pd.Tag] = shop
			}
			// 使用shop信息构造商品链接
			if shop.AffLink != "" {
				switch {
				case strings.Contains(shop.Type, "whmcs"):
					tempList := strings.Split(pd.Url, "&pid=")
					pd.Url = shop.AffLink + "&pid=" + tempList[len(tempList)-1]
				case strings.Contains(shop.Type, "blesta"):
					tempList := strings.Split(shop.AffLink, "/a/")
					pd.Url = pd.Url + "&a=" + tempList[len(tempList)-1]
				case strings.Contains(shop.Type, "hostbill"):
					tempList := strings.Split(pd.Url, "&id=")
					pd.Url = shop.AffLink + "&id=" + tempList[len(tempList)-1]
				}
			}
		}
	}
	return pdProducts, total, err
}

// buildPublicQueryConditions 构建公开查询的条件
// 参数: db *gorm.DB - 数据库查询对象
//
//	info productsReq.ProductsSearch - 搜索条件
//
// 返回: *gorm.DB - 添加了查询条件的数据库查询对象
func (pdService *ProductsService) buildPublicQueryConditions(db *gorm.DB, info productsReq.PublicProductsSearch) *gorm.DB {
	// 字符串字段模糊查询
	stringFields := map[string]string{
		"tag": info.Tag, "cpu": info.Cpu, "memory": info.Memory,
		"disk": info.Disk, "traffic": info.Traffic, "port_speed": info.PortSpeed,
		"location": info.Location, "price": info.Price, "additional": info.Additional,
		"url": info.Url,
	}
	for field, value := range stringFields {
		if value != "" {
			db = db.Where(field+" LIKE ?", "%"+value+"%")
		}
	}
	// 库存数量查询
	if info.Stock != nil {
		db = db.Where("stock > ?", info.Stock)
	} else {
		db = db.Where("cpu <> ?", "")
		db = db.Where("memory <> ?", "")
		db = db.Where("disk <> ?", "")
		db = db.Where("price <> ?", "")
	}
	return db
}
