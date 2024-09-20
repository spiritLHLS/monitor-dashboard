package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type EncryptedLinkSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	ProductId      *int       `json:"product_id" form:"product_id"`
	RedirectUrl    string     `json:"redirectUrl" form:"redirectUrl" `
	ShortCode      string     `json:"shortCode" form:"shortCode"`
	request.PageInfo
}

type BuildEncryptedUrlIds struct {
	Ids []uint `json:"ids"`
}
