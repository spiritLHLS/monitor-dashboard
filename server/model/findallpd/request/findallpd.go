package request

import (
	"server/model/common/request"
	"time"
)

type FindallpdSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Tag            string     `json:"tag" form:"tag" `
	Flag           string     `json:"flag" form:"flag" `
	BillingType    string     `json:"billingType" form:"billingType" `
	PdId           *int       `json:"pdId" form:"pdId" `
	EndId          *int       `json:"endId" form:"endId" `
	Cpu            string     `json:"cpu" form:"cpu" `
	Memory         string     `json:"memory" form:"memory" `
	Disk           string     `json:"disk" form:"disk" `
	Traffic        string     `json:"traffic" form:"traffic" `
	PortSpeed      string     `json:"portSpeed" form:"portSpeed" `
	Location       string     `json:"location" form:"location" `
	Price          string     `json:"price" form:"price" `
	Additional     string     `json:"additional" form:"additional" `
	Url            string     `json:"url" form:"url" `
	OldStock       *int       `json:"oldStock" form:"oldStock" `
	Stock          *int       `json:"stock" form:"stock" `
	MessageId      string     `json:"messageId" form:"messageId" `
	request.PageInfo
}
