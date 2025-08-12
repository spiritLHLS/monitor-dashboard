package request

import (
	"server/model/common/request"
	"time"
)

type InfoSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type InfoTitleSearch struct {
	Title string `json:"title" form:"title" gorm:"column:title;comment:公告标题;"` //标题
}
