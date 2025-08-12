// 自动生成模板Products
package products

import (
	"server/global"
	"time"
)

// products表 结构体  Products
type Products struct {
	global.GVA_MODEL
	Tag           string     `json:"tag" form:"tag" gorm:"column:tag;comment:;size:20;"`                                   //tag字段
	Cpu           string     `json:"cpu" form:"cpu" gorm:"column:cpu;comment:;size:191;"`                                  //cpu字段
	Memory        string     `json:"memory" form:"memory" gorm:"column:memory;comment:;size:191;"`                         //memory字段
	Disk          string     `json:"disk" form:"disk" gorm:"column:disk;comment:;size:191;"`                               //disk字段
	Traffic       string     `json:"traffic" form:"traffic" gorm:"column:traffic;comment:;size:191;"`                      //traffic字段
	PortSpeed     string     `json:"portSpeed" form:"portSpeed" gorm:"column:port_speed;comment:;size:191;"`               //portSpeed字段
	Location      string     `json:"location" form:"location" gorm:"column:location;comment:;size:191;"`                   //location字段
	Price         string     `json:"price" form:"price" gorm:"column:price;comment:;size:191;"`                            //price字段
	Additional    string     `json:"additional" form:"additional" gorm:"column:additional;comment:;type:text;"`            //additional字段
	Url           string     `json:"url" form:"url" gorm:"column:url;comment:;size:191;"`                                  //url字段
	BillingType   string     `json:"billingType" form:"billingType" gorm:"column:billing_type;comment:;size:191;"`         //billingType字段
	PushStock     *int       `json:"pushStock" form:"pushStock" gorm:"column:push_stock;comment:;"`                        // 一对一推送时记录的当前库存
	OldStock      *int       `json:"oldStock" form:"oldStock" gorm:"column:old_stock;comment:;"`                           //oldStock字段
	Stock         *int       `json:"stock" form:"stock" gorm:"column:stock;comment:;"`                                     //stock字段
	MultiCheck    *int       `json:"multiCheck" form:"multiCheck" gorm:"column:multi_check;comment:是否3次检测再推送;"`            //是否3次检测再推送
	Intervals     *int       `json:"intervals" form:"intervals" gorm:"column:intervals;comment:间隔为-1时不进行爬虫;"`              //间隔为-1时不进行爬虫
	MessageId     string     `json:"messageId" form:"messageId" gorm:"column:message_id;comment:自动填充-帖子ID;size:191;"`      //自动填充-帖子ID
	PushIntervals *int       `json:"pushIntervals" form:"pushIntervals" gorm:"column:push_intervals;comment:间隔为-1时不进行推送;"` //间隔为-1时不进行推送
	PushTime      *time.Time `json:"pushTime" form:"pushTime" gorm:"column:push_time;comment:自动填充-最新推送的时间;"`               //自动填充-最新推送的时间
	CreatedBy     *int       `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;"`                     //创建者
	UpdatedBy     *int       `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;"`                     //更新者
	DeletedBy     *int       `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;"`                     //删除者
}

// TableName products表 Products自定义表名 products
func (Products) TableName() string {
	return "products"
}
