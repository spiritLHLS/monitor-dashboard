package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type InfoSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

// TODO 写一个通配搜索的公告接口
type InfoTitleSearch struct {
	Title   string `json:"title" form:"title" gorm:"column:title;comment:公告标题;"`                 //标题
	Content string `json:"content" form:"content" gorm:"column:content;comment:公告内容;type:text;"` //内容
	UserID  *int   `json:"userID" form:"userID" gorm:"column:user_id;comment:发布者;"`              //作者
}
