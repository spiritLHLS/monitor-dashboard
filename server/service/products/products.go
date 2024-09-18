package products

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/products"
	productsReq "github.com/flipped-aurora/gin-vue-admin/server/model/products/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shops"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cryptourl/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cryptourl/service"
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
func (pdService *ProductsService) GetProductsPublic(info productsReq.PublicProductsSearch) ([]products.Products, int64, error) {
	db := global.GVA_DB.Model(&products.Products{})
	db = pdService.buildPublicQueryConditions(db, info)
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if info.PageSize > 0 {
		db = db.Limit(info.PageSize).Offset(info.PageSize * (info.Page - 1))
	}
	var pdProducts []products.Products
	if err := db.Find(&pdProducts).Error; err != nil {
		return nil, 0, err
	}
	if total > 0 {
		if err := pdService.processProducts(&pdProducts); err != nil {
			return nil, 0, err
		}
	}
	return pdProducts, total, nil
}

// processProducts 处理商品信息，包括更新URL和创建短链接
func (pdService *ProductsService) processProducts(pdProducts *[]products.Products) error {
	// 创建商店缓存，用于存储已查询的商店信息
	shopCache := make(map[string]*shops.Shops)
	// 创建待处理URL列表
	urlsToProcess := make([]string, 0, len(*pdProducts))
	// 创建URL到商品的映射，用于快速查找
	urlToProductMap := make(map[string]*products.Products)
	// 第一次遍历：收集所有需要处理的URL并更新商品URL
	for i := range *pdProducts {
		pd := &(*pdProducts)[i]
		// 获取商店信息
		shop, err := pdService.getShopInfo(pd.Tag, shopCache)
		if err != nil {
			continue
		}
		// 如果商店有联盟链接，更新商品URL
		if shop.AffLink != "" {
			pdService.updateProductUrl(pd, shop)
		}
		// 将更新后的URL添加到待处理列表
		urlsToProcess = append(urlsToProcess, pd.Url)
		// 建立URL到商品的映射
		urlToProductMap[pd.Url] = pd
	}
	// 批量获取或创建加密链接
	encryptedLinks, err := pdService.batchGetOrCreateEncryptedLinks(urlsToProcess)
	if err != nil {
		return err
	}
	// 第二次遍历：更新商品的短链接
	for _, el := range encryptedLinks {
		if pd, ok := urlToProductMap[el.RedirectUrl]; ok {
			pd.Url = el.ShortCode
		}
	}
	return nil
}

// getShopInfo 获取商店信息，使用缓存减少数据库查询
func (pdService *ProductsService) getShopInfo(tag string, shopCache map[string]*shops.Shops) (*shops.Shops, error) {
	// 如果缓存中存在，直接返回
	if shop, exists := shopCache[tag]; exists {
		return shop, nil
	}
	// 如果缓存中不存在，从数据库查询
	shop := &shops.Shops{}
	if err := global.GVA_DB.Where("tag = ?", tag).First(shop).Error; err != nil {
		return nil, err
	}
	// 将查询结果存入缓存
	shopCache[tag] = shop
	return shop, nil
}

// updateProductUrl 根据商店类型更新商品URL
func (pdService *ProductsService) updateProductUrl(pd *products.Products, shop *shops.Shops) {
	switch {
	case strings.Contains(shop.Type, "whmcs"):
		pd.Url = shop.AffLink + "&pid=" + strings.Split(pd.Url, "&pid=")[1]
	case strings.Contains(shop.Type, "blesta"):
		pd.Url += "&a=" + strings.Split(shop.AffLink, "/a/")[1]
	case strings.Contains(shop.Type, "hostbill"):
		pd.Url = shop.AffLink + "&id=" + strings.Split(pd.Url, "&id=")[1]
	}
}

// batchGetOrCreateEncryptedLinks 批量获取或创建加密链接
func (pdService *ProductsService) batchGetOrCreateEncryptedLinks(urls []string) ([]model.EncryptedLink, error) {
	var existingLinks []model.EncryptedLink
	// 批量查询已存在的加密链接
	err := global.GVA_DB.Where("redirect_url IN ?", urls).Find(&existingLinks).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	// 创建已存在URL的映射，用于快速查找
	existingUrlMap := make(map[string]bool)
	for _, link := range existingLinks {
		existingUrlMap[link.RedirectUrl] = true
	}
	// 收集并创建新的加密链接
	for _, url := range urls {
		if !existingUrlMap[url] {
			newLink := model.EncryptedLink{RedirectUrl: url}
			if err := service.EncryptedLink.CreateEncryptedLink(&newLink); err != nil {
				return nil, err
			}
			existingLinks = append(existingLinks, newLink)
		}
	}
	return existingLinks, nil
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
		db = db.Where("price <> ? AND price IS NOT NULL", "").Order("LENGTH(tag)")
	}
	return db
}
