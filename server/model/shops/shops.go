// 自动生成模板Shops
package shops

import (
	"server/global"
)

// shops表 结构体  Shops
type Shops struct {
	global.GVA_MODEL
	Tag            string `json:"tag" form:"tag" gorm:"column:tag;comment:;size:500;"`                                    //频道TAG
	Mid            string `json:"mid" form:"mid" gorm:"column:mid;comment:;size:191;"`                                    //消息编号
	Type           string `json:"type" form:"type" gorm:"column:type;comment:;size:191;"`                                 //商家类型
	AffLink        string `json:"affLink" form:"affLink" gorm:"column:aff_link;comment:;size:191;"`                       //推广链接
	ShopLink       string `json:"shopLink" form:"shopLink" gorm:"column:shop_link;comment:;size:191;"`                    //商家链接
	AdditionalTags string `json:"additionalTags" form:"additionalTags" gorm:"column:additional_tags;comment:;size:1000;"` //额外标签
	CreatedBy      *int   `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:20;"`               //创建者
	UpdatedBy      *int   `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:20;"`               //更新者
	DeletedBy      *int   `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;size:20;"`               //删除者
}

// TableName shops表 Shops自定义表名 shops
func (Shops) TableName() string {
	return "shops"
}
