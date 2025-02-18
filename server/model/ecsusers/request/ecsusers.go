package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/google/uuid"
	"time"
)

type EcsUsersSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	UUID           uuid.UUID  `json:"uuid" form:"uuid" `
	Username       string     `json:"username" form:"username" `
	Nickname       string     `json:"nickname" form:"nickname" `
	Password       string     `json:"password" form:"password" `
	IsFrozen       *bool      `json:"isFrozen" form:"isFrozen" `
	PushChannel1   string     `json:"pushChannel1" form:"pushChannel1" `
	PushChannel2   string     `json:"pushChannel2" form:"pushChannel2" `
	PushChannel3   string     `json:"pushChannel3" form:"pushChannel3" `
	TGID           string     `json:"tgID" form:"tgID" `
	QQNumber       string     `json:"qqNumber" form:"qqNumber" `
	WeChatNumber   string     `json:"weChatNumber" form:"weChatNumber" `
	Email          string     `json:"email" form:"email" `
	Additional     string     `json:"additional" form:"additional" `
	StartLevel     *int       `json:"startLevel" form:"startLevel"`
	EndLevel       *int       `json:"endLevel" form:"endLevel"`
	request.PageInfo
}

type AdminChangePasswordReq struct {
	UserID   uint   `json:"userID"`
	Password string `json:"password"`
}
