package model

import (
	"server/global"
)

// EncryptedLink 加密链接 结构体
type EncryptedLink struct {
	global.GVA_MODEL
	ProductId   *uint  `json:"product_id" form:"product_id" gorm:"column:product_id;comment:商品ID;"`                 //商品ID
	RedirectUrl string `json:"redirectUrl" form:"redirectUrl" gorm:"column:redirect_url;unique;comment:需要跳转的目标链接;"` //跳转链接
	ShortCode   string `json:"shortCode" form:"shortCode" gorm:"column:short_code;comment:短代码用于前端请求;"`
}

// TableName 加密链接 EncryptedLink自定义表名 EncryptedLinks
func (EncryptedLink) TableName() string {
	return "EncryptedLinks"
}
