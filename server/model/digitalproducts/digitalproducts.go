// 自动生成模板DigitalProducts
package digitalproducts

import (
	"server/global"
)

// 数字商品表 结构体  DigitalProducts
type DigitalProducts struct {
	global.GVA_MODEL
	Tag        *string `json:"tag" form:"tag" gorm:"column:tag;comment:;size:20;"`                              //TAG
	Cpu        *int    `json:"cpu" form:"cpu" gorm:"column:cpu;comment:核心数;size:20;"`                           //核心
	Memory     *int    `json:"memory" form:"memory" gorm:"column:memory;comment:以GB计算的大小;size:20;"`             //内存
	Disk       *int    `json:"disk" form:"disk" gorm:"column:disk;comment:以GB计算的大小;size:20;"`                   //硬盘
	Traffic    *int    `json:"traffic" form:"traffic" gorm:"column:traffic;comment:以T计算的大小;size:20;"`           //流量
	PortSpeed  *int    `json:"portSpeed" form:"portSpeed" gorm:"column:port_speed;comment:以Gbps计算的大小;size:20;"` //带宽
	Location   *string `json:"location" form:"location" gorm:"column:location;comment:;size:191;"`              //位置
	Price      *int    `json:"price" form:"price" gorm:"column:price;comment:价格大小;size:20;"`                    //价格
	PriceUnit  *string `json:"priceUnit" form:"priceUnit" gorm:"column:price_unit;comment:价格单位;size:20;"`       //价格单位
	Additional *string `json:"additional" form:"additional" gorm:"column:additional;comment:;size:191;"`        //其他
	Stock      *int    `json:"stock" form:"stock" gorm:"column:stock;comment:库存存在但无数显示为10;size:20;"`            //库存
	OriginId   *int    `json:"originId" form:"originId" gorm:"column:origin_id;comment:原表主键ID;size:20;"`        //原表ID
	CreatedBy  uint    `gorm:"column:created_by;comment:创建者"`
	UpdatedBy  uint    `gorm:"column:updated_by;comment:更新者"`
	DeletedBy  uint    `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 数字商品表 DigitalProducts自定义表名 digitalproducts
func (DigitalProducts) TableName() string {
	return "digitalproducts"
}
