package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type TgchannelSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Channel        string     `json:"channel" form:"channel" `
	Tgid           string     `json:"tgid" form:"tgid" `
	Flag           string     `json:"flag" form:"flag" `
	AdditionalKey  string     `json:"additionalKey" form:"additionalKey" `
	Tokens         string     `json:"tokens" form:"tokens" `
	Tags           string     `json:"tags" form:"tags" `
	request.PageInfo
}
