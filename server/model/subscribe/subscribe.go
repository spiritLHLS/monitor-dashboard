// 自动生成模板Subscribe
package subscribe

import (
	"github.com/google/uuid"
	"server/global"
	"time"
)

// 订阅 结构体  Subscribe
type Subscribe struct {
	global.GVA_MODEL
	UserUuid      uuid.UUID  `json:"user_uuid" form:"user_uuid" gorm:"primarykey;column:user_uuid;comment:用户ID;"`     //用户ID
	ProductId     *int       `json:"product_id" form:"product_id" gorm:"column:product_id;comment:商品ID;"`             //商品ID
	Status        *int       `json:"status" form:"status" gorm:"column:status;comment:订阅状态;"`                         //订阅状态
	NotifyChannel string     `json:"notify_channel" form:"notify_channel" gorm:"column:notify_channel;comment:通知渠道;"` //通知渠道
	LastUpdate    *time.Time `json:"last_update" form:"last_update" gorm:"column:last_update;comment:最后更新时间;"`        //最后更新
}

// TableName 订阅 Subscribe自定义表名 subscribe
func (Subscribe) TableName() string {
	return "subscribe"
}

// UpdateStatusIDs 更新订阅的状态表
type UpdateStatusIDs struct {
	IDs []*int `json:"IDs"`
}
