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
func (subService *SubscribeService) CreateSubscribe(sub *subscribe.Subscribe) (err error) {
	return global.GVA_DB.Create(sub).Error
}

// DeleteSubscribe 删除订阅记录
func (subService *SubscribeService) DeleteSubscribe(ID string) (err error) {
	return global.GVA_DB.Delete(&subscribe.Subscribe{}, "id = ?", ID).Error
}

// DeleteSubscribeByIds 批量删除订阅记录
func (subService *SubscribeService) DeleteSubscribeByIds(IDs []string) (err error) {
	return global.GVA_DB.Delete(&[]subscribe.Subscribe{}, "id in ?", IDs).Error
}

// UpdateSubscribe 更新订阅记录
func (subService *SubscribeService) UpdateSubscribe(sub subscribe.Subscribe) (err error) {
	return global.GVA_DB.Model(&subscribe.Subscribe{}).Where("id = ?", sub.ID).Updates(&sub).Error
}

// GetSubscribe 根据ID获取订阅记录
func (subService *SubscribeService) GetSubscribe(ID string) (sub subscribe.Subscribe, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&sub).Error
	return
}

// GetSubscribeInfoList 分页获取订阅记录
func (subService *SubscribeService) GetSubscribeInfoList(info subscribeReq.SubscribeSearch) (list []subscribe.Subscribe, total int64, err error) {
	db := global.GVA_DB.Model(&subscribe.Subscribe{})

	// 优化：使用map来存储查询条件，减少重复代码
	conditions := map[string]interface{}{
		"user_uuid = ?":         info.UserUuid,
		"product_id = ?":        info.ProductId,
		"status = ?":            info.Status,
		"notify_channel LIKE ?": "%" + info.NotifyChannel + "%",
	}

	for condition, value := range conditions {
		if value != nil && value != "" {
			db = db.Where(condition, value)
		}
	}

	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	if info.LastUpdate != nil {
		db = db.Where("last_update < ?", info.LastUpdate)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if info.PageSize != 0 {
		db = db.Limit(info.PageSize).Offset(info.PageSize * (info.Page - 1))
	}

	err = db.Find(&list).Error
	return
}

// buildQueryConditions 构建查询条件
func (subService *SubscribeService) buildQueryConditions(db *gorm.DB, info interface{}) *gorm.DB {
	var stringFields map[string]string
	var oldStock, stock *int
	// 使用类型断言确定实际的类型
	switch v := info.(type) {
	case productsReq.ProductsSearch:
		stringFields = map[string]string{
			"tag": v.Tag, "cpu": v.Cpu, "memory": v.Memory,
			"disk": v.Disk, "traffic": v.Traffic, "port_speed": v.PortSpeed,
			"location": v.Location, "price": v.Price, "additional": v.Additional,
		}
		oldStock = v.OldStock
		stock = v.Stock
	case subscribeReq.SubscribeSearch:
		stringFields = map[string]string{
			"tag": v.Tag, "cpu": v.Cpu, "memory": v.Memory,
			"disk": v.Disk, "traffic": v.Traffic, "port_speed": v.PortSpeed,
			"location": v.Location, "price": v.Price, "additional": v.Additional,
		}
		oldStock = v.OldStock
		stock = v.Stock
	default:
		// 如果既不是 ProductsSearch 也不是 SubscribeSearch，则返回未修改的 db
		return db
	}
	// 字符串字段模糊查询
	for field, value := range stringFields {
		if value != "" {
			db = db.Where(field+" LIKE ?", "%"+value+"%")
		}
	}
	// 数值字段精确查询
	if oldStock != nil {
		db = db.Where("old_stock > ?", oldStock)
	}
	if stock != nil {
		db = db.Where("stock > ?", stock)
	}
	// 非空查询
	//db = db.Where("cpu <> ?", "")
	//db = db.Where("memory <> ?", "")
	//db = db.Where("disk <> ?", "")
	db = db.Where("price <> ? AND price IS NOT NULL", "").Order("LENGTH(tag)")
	return db
}

// SelfGetSub 前端用户获取自己已订阅的商品
func (subService *SubscribeService) SelfGetSub(uuid uuid.UUID, info subscribeReq.SubscribeSearch) (list []products.SubProducts, total int64, err error) {
	var subs []subscribe.Subscribe
	if err = global.GVA_DB.Where("user_uuid = ?", uuid).Find(&subs).Error; err != nil {
		return nil, 0, err
	}
	if len(subs) == 0 {
		return []products.SubProducts{}, 0, fmt.Errorf("subscription not found")
	}
	productIDs := make([]int, len(subs))
	for i, sub := range subs {
		productIDs[i] = *sub.ProductId
	}
	db := global.GVA_DB.Model(&products.Products{}).Where("id IN ?", productIDs)
	db = subService.buildQueryConditions(db, info)
	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	if info.PageSize > 0 {
		db = db.Limit(info.PageSize).Offset(info.PageSize * (info.Page - 1))
	}
	err = db.Find(&list).Error
	return
}

// SelfCreateSub 前端用户创建关联的商品推送记录
func (subService *SubscribeService) SelfCreateSub(uuid uuid.UUID, pdid int, notifyChannel string) (err error) {
	var existingSub subscribe.Subscribe
	if err := global.GVA_DB.Where("user_uuid = ? AND product_id = ?", uuid, pdid).First(&existingSub).Error; err == nil {
		return fmt.Errorf("subscription already exists")
	} else if err != gorm.ErrRecordNotFound {
		return err
	}
	now := time.Now()
	status := 0
	newSub := subscribe.Subscribe{
		UserUuid:      uuid,
		ProductId:     &pdid,
		Status:        &status,
		NotifyChannel: notifyChannel,
		LastUpdate:    &now,
	}
	return global.GVA_DB.Create(&newSub).Error
}

// SelfDeleteSub 前端用户删除自己已订阅的商品
func (subService *SubscribeService) SelfDeleteSub(uuid uuid.UUID, pdid int) (err error) {
	return global.GVA_DB.Where("user_uuid = ? AND product_id = ?", uuid, pdid).Delete(&subscribe.Subscribe{}).Error
}

// SelfUpdateSub 前端用户更新自己已订阅的商品信息
func (subService *SubscribeService) SelfUpdateSub(uuid uuid.UUID, pdid int, notifyChannel string) (err error) {
	result := global.GVA_DB.Model(&subscribe.Subscribe{}).
		Where("user_uuid = ? AND product_id = ?", uuid, pdid).
		Update("notify_channel", notifyChannel)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("subscription not found")
	}
	return nil
}

// SelfGetAllPd 前端用户获取所有的商品
func (subService *SubscribeService) SelfGetAllPd(info productsReq.ProductsSearch) (list []products.SubProducts, total int64, err error) {
	db := global.GVA_DB.Model(&products.Products{})
	db = subService.buildQueryConditions(db, info)
	if err = db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if info.PageSize > 0 {
		db = db.Limit(info.PageSize).Offset(info.PageSize * (info.Page - 1))
	}
	err = db.Find(&list).Error
	return
}
