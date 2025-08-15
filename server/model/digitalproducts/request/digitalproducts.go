package request

import (
	"server/model/common/request"
	"time"
)

type DigitalProductsSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Tag            *string    `json:"tag" form:"tag"`
	Cpu            *float64   `json:"cpu" form:"cpu"`
	Memory         *float64   `json:"memory" form:"memory"`
	Disk           *float64   `json:"disk" form:"disk"`
	Traffic        *float64   `json:"traffic" form:"traffic"`
	PortSpeed      *float64   `json:"portSpeed" form:"portSpeed"`
	Location       *string    `json:"location" form:"location"`
	Price          *float64   `json:"price" form:"price"`
	PriceUnit      *string    `json:"priceUnit" form:"priceUnit"`
	request.PageInfo
}
