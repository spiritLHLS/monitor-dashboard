package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/gofrs/uuid/v5"
	"time"
)

type SubscribeSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	UserUuid       uuid.UUID  `json:"user_uuid" form:"user_uuid" `
	ProductId      *int       `json:"product_id" form:"product_id" `
	Status         *int       `json:"status" form:"status" `
	NotifyChannel  string     `json:"notify_channel" form:"notify_channel" `
	LastUpdate     *time.Time `json:"last_update" form:"last_update" `
	request.PageInfo
}
