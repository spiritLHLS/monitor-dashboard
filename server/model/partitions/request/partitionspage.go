package request

import (
	"server/model/common/request"
	"time"
)

type PartitionspageSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	TgTag          string     `json:"tg_tag" form:"tg_tag" `
	Name           string     `json:"name" form:"name" `
	Link           string     `json:"link" form:"link" `
	Type           string     `json:"type" form:"type" `
	Num            *int       `json:"num" form:"num" `
	Flag           string     `json:"flag" form:"flag" `
	Additional     string     `json:"additional" form:"additional" `
	Intervals      *int       `json:"intervals" form:"intervals" `
	request.PageInfo
}
