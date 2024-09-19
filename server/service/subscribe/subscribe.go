package subscribe

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/products"
	productsReq "github.com/flipped-aurora/gin-vue-admin/server/model/products/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/subscribe"
	subscribeReq "github.com/flipped-aurora/gin-vue-admin/server/model/subscribe/request"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
	"time"
)

type SubscribeService struct{}

// CreateSubscribe 创建订阅记录
// Author [piexlmax](https://github.com/piexlmax)
func (subService *SubscribeService) CreateSubscribe(sub *subscribe.Subscribe) (err error) {
	err = global.GVA_DB.Create(sub).Error
	return err
}

// DeleteSubscribe 删除订阅记录
// Author [piexlmax](https://github.com/piexlmax)
func (subService *SubscribeService) DeleteSubscribe(ID string) (err error) {
	err = global.GVA_DB.Delete(&subscribe.Subscribe{}, "id = ?", ID).Error
	return err
}

// DeleteSubscribeByIds 批量删除订阅记录
// Author [piexlmax](https://github.com/piexlmax)
func (subService *SubscribeService) DeleteSubscribeByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]subscribe.Subscribe{}, "id in ?", IDs).Error
	return err
}

// UpdateSubscribe 更新订阅记录
// Author [piexlmax](https://github.com/piexlmax)
func (subService *SubscribeService) UpdateSubscribe(sub subscribe.Subscribe) (err error) {
	err = global.GVA_DB.Model(&subscribe.Subscribe{}).Where("id = ?", sub.ID).Updates(&sub).Error
	return err
}

// GetSubscribe 根据ID获取订阅记录
// Author [piexlmax](https://github.com/piexlmax)
func (subService *SubscribeService) GetSubscribe(ID string) (sub subscribe.Subscribe, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&sub).Error
	return
}

// GetSubscribeInfoList 分页获取订阅记录
// Author [piexlmax](https://github.com/piexlmax)
func (subService *SubscribeService) GetSubscribeInfoList(info subscribeReq.SubscribeSearch) (list []subscribe.Subscribe, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&subscribe.Subscribe{})
	var subs []subscribe.Subscribe
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.UserUuid.String() != "" {
		db = db.Where("user_uuid = ?", info.UserUuid.String())
	}
	if info.ProductId != nil {
		db = db.Where("product_id = ?", info.ProductId)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.NotifyChannel != "" {
		db = db.Where("notify_channel LIKE ?", "%"+info.NotifyChannel+"%")
	}
	if info.LastUpdate != nil {
		db = db.Where("last_update < ?", info.LastUpdate)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&subs).Error
	return subs, total, err
}

// SelfGetSub 前端用户获取自己已订阅的商品
// Author [yourname](https://github.com/yourname)
func (subService *SubscribeService) SelfGetSub(uuid uuid.UUID) ([]subscribe.Subscribe, error) {
	var subs []subscribe.Subscribe
	err := global.GVA_DB.Model(&subscribe.Subscribe{}).Where("user_uuid = ?", uuid.String()).Find(&subs).Error
	return subs, err
}

// SelfCreateSub 前端用户创建关联的商品推送记录
// Author [yourname](https://github.com/yourname)
func (subService *SubscribeService) SelfCreateSub(uuid uuid.UUID, pdid int, notifyChannel string) (err error) {
	var existingSub subscribe.Subscribe
	result := global.GVA_DB.Model(&subscribe.Subscribe{}).Where("user_uuid = ? AND product_id = ?", uuid.String(), pdid).First(&existingSub)
	if result.Error == nil {
		// 订阅已存在，返回错误或者可以选择静默失败
		return fmt.Errorf("subscription already exists")
	}
	if result.Error != gorm.ErrRecordNotFound {
		// 处理其他查询错误
		return result.Error
	}
	// 订阅不存在，创建新的订阅
	now := time.Now()
	status := 0
	newSub := subscribe.Subscribe{
		UserUuid:      uuid,
		ProductId:     &pdid,
		Status:        &status,
		NotifyChannel: notifyChannel,
		LastUpdate:    &now,
	}
	err = global.GVA_DB.Create(&newSub).Error
	return err
}

// SelfDeleteSub 前端用户删除自己已订阅的商品
// Author [yourname](https://github.com/yourname)
func (subService *SubscribeService) SelfDeleteSub(uuid uuid.UUID, pdid int) (err error) {
	err = global.GVA_DB.Model(&subscribe.Subscribe{}).Where("user_uuid = ? AND product_id = ?", uuid.String(), pdid).Delete(&subscribe.Subscribe{}).Error
	return err
}

// SelfUpdateSub 前端用户更新自己已订阅的商品信息
// Author [yourname](https://github.com/yourname)
func (subService *SubscribeService) SelfUpdateSub(uuid uuid.UUID, pdid int, notifyChannel string) (err error) {
	var existingSub subscribe.Subscribe
	result := global.GVA_DB.Model(&subscribe.Subscribe{}).Where("user_uuid = ? AND product_id = ?", uuid.String(), pdid).First(&existingSub)
	if result.Error == gorm.ErrRecordNotFound {
		// 订阅不存在，返回错误
		return fmt.Errorf("subscription not found")
	}
	if result.Error != nil {
		// 处理其他查询错误
		return result.Error
	}
	// 订阅存在，更新它
	updates := map[string]interface{}{
		"notify_channel": notifyChannel,
	}
	err = global.GVA_DB.Model(&existingSub).Updates(updates).Error
	return err
}

// buildQueryConditions 构建查询条件
// 参数: db *gorm.DB - 数据库查询对象
//
//	info productsReq.ProductsSearch - 搜索条件
//
// 返回: *gorm.DB - 添加了查询条件的数据库查询对象
func (subService *SubscribeService) buildQueryConditions(db *gorm.DB, info productsReq.ProductsSearch) *gorm.DB {
	// 时间范围查询
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	// 字符串字段模糊查询
	stringFields := map[string]string{
		"tag": info.Tag, "cpu": info.Cpu, "memory": info.Memory,
		"disk": info.Disk, "traffic": info.Traffic, "port_speed": info.PortSpeed,
		"location": info.Location, "price": info.Price, "additional": info.Additional,
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
	return db
}

// SelfGetAllPd 前端用户获取所有的商品
// Author [yourname](https://github.com/yourname)
func (subService *SubscribeService) SelfGetAllPd(info productsReq.ProductsSearch) (list []products.SubProducts, total int64, err error) {
	// 计算分页参数
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建数据库查询对象
	db := global.GVA_DB.Model(&products.Products{})
	// 构建查询条件
	db = subService.buildQueryConditions(db, info)
	// 获取总记录数
	if err = db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	// 应用分页
	if limit > 0 {
		db = db.Limit(limit).Offset(offset)
	}
	// 执行查询
	var pdProducts []products.SubProducts
	err = db.Find(&pdProducts).Error
	return pdProducts, total, err
}
