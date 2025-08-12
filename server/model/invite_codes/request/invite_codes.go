package request

import (
	"server/model/common/request"
	"time"
)

type InviteCodesSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Code           string     `json:"code" form:"code" `
	Status         *int       `json:"status" form:"status" `
	Expires_at     *time.Time `json:"expires_at" form:"expires_at" `
	User_uuid      string     `json:"user_uuid" form:"user_uuid" `
	Remarks        string     `json:"remarks" form:"remarks" `
	request.PageInfo
}
