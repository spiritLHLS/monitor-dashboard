package subscribe

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/subscribe"
	subscribeReq "github.com/flipped-aurora/gin-vue-admin/server/model/subscribe/request"
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
	if info.UserUuid != nil {
		db = db.Where("user_uuid = ?", info.UserUuid)
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
func (subService *SubscribeService) SelfGetSub() {
}

// SelfCreateSub 前端用户创建关联的商品推送记录
// Author [yourname](https://github.com/yourname)
func (subService *SubscribeService) SelfCreateSub() (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&subscribe.Subscribe{})
	return db.Error
}

// SelfDeleteSub 前端用户删除自己已订阅的商品
// Author [yourname](https://github.com/yourname)
func (subService *SubscribeService) SelfDeleteSub() (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&subscribe.Subscribe{})
	return db.Error
}

// SelfGetAllPd 前端用户获取所有的商品
// Author [yourname](https://github.com/yourname)
func (subService *SubscribeService)SelfGetAllPd() (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&subscribe.Subscribe{})
    return db.Error
}

