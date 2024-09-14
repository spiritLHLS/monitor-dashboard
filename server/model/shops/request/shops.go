package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ShopsSearch struct {
	Tag            string     `json:"tag" form:"tag" `
	Mid            string     `json:"mid" form:"mid" `
	Type           string     `json:"type" form:"type" `
	AffLink        string     `json:"affLink" form:"affLink" `
	ShopLink       string     `json:"shopLink" form:"shopLink" `
	AdditionalTags string     `json:"additionalTags" form:"additionalTags" `
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
