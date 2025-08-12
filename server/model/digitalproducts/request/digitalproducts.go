package request

import (
	"server/model/common/request"
	"time"
)

type DigitalProductsSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	PriceUnit      *string    `json:"priceUnit" form:"priceUnit" `
	request.PageInfo
}
