package request

import (
	"server/model/common/request"
	"time"
)

type PusherConfigSearch struct {
	StartCreatedAt       *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt         *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	PushType             string     `json:"pushType" form:"pushType" `
	ConfigName           string     `json:"configName" form:"configName" `
	ConfigValue          string     `json:"configValue" form:"configValue" `
	MaxCharactersPerPush *int       `json:"maxCharactersPerPush" form:"maxCharactersPerPush" `
	LastPushTime         *time.Time `json:"lastPushTime" form:"lastPushTime" `
	IsActive             *bool      `json:"isActive" form:"isActive" `
	request.PageInfo
}
