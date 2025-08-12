package request

import (
	"server/model/common/request"
	"time"
)

type ProductsSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Tag           string     `json:"tag" form:"tag" `
	Cpu           string     `json:"cpu" form:"cpu" `
	Memory        string     `json:"memory" form:"memory" `
	Disk          string     `json:"disk" form:"disk" `
	Traffic       string     `json:"traffic" form:"traffic" `
	PortSpeed     string     `json:"portSpeed" form:"portSpeed" `
	Location      string     `json:"location" form:"location" `
	Price         string     `json:"price" form:"price" `
	Additional    string     `json:"additional" form:"additional" `
	Url           string     `json:"url" form:"url" `
	BillingType   string     `json:"billingType" form:"billingType" `
	PushStock     *int       `json:"pushStock" form:"pushStock"`
	OldStock      *int       `json:"oldStock" form:"oldStock" `
	Stock         *int       `json:"stock" form:"stock" `
	MultiCheck    *int       `json:"multiCheck" form:"multiCheck" `
	Intervals     *int       `json:"intervals" form:"intervals" `
	MessageId     string     `json:"messageId" form:"messageId" `
	PushIntervals *int       `json:"pushIntervals" form:"pushIntervals" `
	PushTime      *time.Time `json:"pushTime" form:"pushTime" `
	request.PageInfo
}

type PublicProductsSearch struct {
	Tag        string `json:"tag" form:"tag" `
	Cpu        string `json:"cpu" form:"cpu" `
	Memory     string `json:"memory" form:"memory" `
	Disk       string `json:"disk" form:"disk" `
	Traffic    string `json:"traffic" form:"traffic" `
	PortSpeed  string `json:"portSpeed" form:"portSpeed" `
	Location   string `json:"location" form:"location" `
	Price      string `json:"price" form:"price" `
	Additional string `json:"additional" form:"additional" `
	Url        string `json:"url" form:"url" `
	Stock      *int   `json:"stock" form:"stock" `
	request.PageInfo
}
