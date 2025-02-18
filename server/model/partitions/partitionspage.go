// 自动生成模板Partitionspage
package partitions

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// partitionspage表 结构体  Partitionspage
type Partitionspage struct {
	global.GVA_MODEL
	TgTag      string `json:"tgTag" form:"tgTag" gorm:"column:tg_tag;comment:;size:191;"`                         //Tag
	Name       string `json:"name" form:"name" gorm:"column:name;comment:;size:191;"`                             //分区名字
	Link       string `json:"link" form:"link" gorm:"column:link;comment:;size:191;"`                             //分区链接
	Type       string `json:"type" form:"type" gorm:"column:type;comment:;size:191;"`                             //分区类型
	Num        *int   `json:"num" form:"num" gorm:"column:num;comment:;size:19;"`                                 //识别数量
	Flag       string `json:"flag" form:"flag" gorm:"column:flag;comment:防御盾验证类型;"`                               //防御盾验证类型
	Additional string `json:"additional" form:"additional" gorm:"column:additional;comment:;size:191;type:text;"` //其他信息
	Intervals  *int   `json:"intervals" form:"intervals" gorm:"column:intervals;comment:;size:19;"`               //爬虫间隔
	CreatedBy  *int   `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:20;"`           //创建者
	UpdatedBy  *int   `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:20;"`           //更新者
	DeletedBy  *int   `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;size:20;"`           //删除者
}

// TableName partitionspage表 Partitionspage自定义表名 partitionspage
func (Partitionspage) TableName() string {
	return "partitionspage"
}
