package request

import (
	"github.com/google/uuid"
	"server/global"
	productsReq "server/model/products/request"
	"time"
)

type SubscribeSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	UserUuid       uuid.UUID  `json:"user_uuid" form:"user_uuid" `
	ProductId      *int       `json:"product_id" form:"product_id" `
	Status         *int       `json:"status" form:"status" `
	NotifyChannel  string     `json:"notify_channel" form:"notify_channel" `
	LastUpdate     *time.Time `json:"last_update" form:"last_update" `
	productsReq.ProductsSearch
}

type SubProducts struct {
	global.GVA_MODEL
	Tag        string   `json:"tag" form:"tag" gorm:"column:tag;comment:;size:20;"`                        //tag字段
	Cpu        *float64 `json:"cpu" form:"cpu" gorm:"column:cpu;comment:;"`                                //cpu字段
	Memory     *float64 `json:"memory" form:"memory" gorm:"column:memory;comment:;"`                       //memory字段
	Disk       *float64 `json:"disk" form:"disk" gorm:"column:disk;comment:;"`                             //disk字段
	Traffic    *float64 `json:"traffic" form:"traffic" gorm:"column:traffic;comment:;"`                    //traffic字段
	PortSpeed  *float64 `json:"portSpeed" form:"portSpeed" gorm:"column:port_speed;comment:;"`             //portSpeed字段
	Location   string   `json:"location" form:"location" gorm:"column:location;comment:;size:191;"`        //location字段
	Price      *float64 `json:"price" form:"price" gorm:"column:price;comment:;"`                          //price字段
	Additional string   `json:"additional" form:"additional" gorm:"column:additional;comment:;type:text;"` //additional字段
	OldStock   *int     `json:"oldStock" form:"oldStock" gorm:"column:old_stock;comment:;"`                //oldStock字段
	Stock      *int     `json:"stock" form:"stock" gorm:"column:stock;comment:;"`                          //stock字段
}

type SubProductsWithNotify struct {
	global.GVA_MODEL
	Tag           string   `json:"tag" form:"tag" gorm:"column:tag;comment:;size:20;"`                              //tag字段
	Cpu           *float64 `json:"cpu" form:"cpu" gorm:"column:cpu;comment:;"`                                      //cpu字段
	Memory        *float64 `json:"memory" form:"memory" gorm:"column:memory;comment:;"`                             //memory字段
	Disk          *float64 `json:"disk" form:"disk" gorm:"column:disk;comment:;"`                                   //disk字段
	Traffic       *float64 `json:"traffic" form:"traffic" gorm:"column:traffic;comment:;"`                          //traffic字段
	PortSpeed     *float64 `json:"portSpeed" form:"portSpeed" gorm:"column:port_speed;comment:;"`                   //portSpeed字段
	Location      string   `json:"location" form:"location" gorm:"column:location;comment:;size:191;"`              //location字段
	Price         *float64 `json:"price" form:"price" gorm:"column:price;comment:;"`                                //price字段
	Additional    string   `json:"additional" form:"additional" gorm:"column:additional;comment:;type:text;"`       //additional字段
	OldStock      *int     `json:"oldStock" form:"oldStock" gorm:"column:old_stock;comment:;"`                      //oldStock字段
	Stock         *int     `json:"stock" form:"stock" gorm:"column:stock;comment:;"`                                //stock字段
	Status        *int     `json:"status" form:"status" `                                                           // status字段
	NotifyChannel string   `json:"notify_channel" form:"notify_channel" gorm:"column:notify_channel;comment:通知渠道;"` //通知渠道
}

// SubProductsSearch 订阅商品搜索结构体，支持范围搜索
type SubProductsSearch struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`

	// 字符串字段 - 模糊搜索
	Tag        string `json:"tag" form:"tag"`
	Location   string `json:"location" form:"location"`
	Additional string `json:"additional" form:"additional"`

	// 数值字段 - 范围搜索
	CpuMin       *float64 `json:"cpuMin" form:"cpuMin"`
	CpuMax       *float64 `json:"cpuMax" form:"cpuMax"`
	MemoryMin    *float64 `json:"memoryMin" form:"memoryMin"`
	MemoryMax    *float64 `json:"memoryMax" form:"memoryMax"`
	DiskMin      *float64 `json:"diskMin" form:"diskMin"`
	DiskMax      *float64 `json:"diskMax" form:"diskMax"`
	TrafficMin   *float64 `json:"trafficMin" form:"trafficMin"`
	TrafficMax   *float64 `json:"trafficMax" form:"trafficMax"`
	PortSpeedMin *float64 `json:"portSpeedMin" form:"portSpeedMin"`
	PortSpeedMax *float64 `json:"portSpeedMax" form:"portSpeedMax"`
	PriceMin     *float64 `json:"priceMin" form:"priceMin"`
	PriceMax     *float64 `json:"priceMax" form:"priceMax"`
	Stock        *int     `json:"stock" form:"stock"`

	// 排序字段
	SortBy    string `json:"sortBy" form:"sortBy"`
	SortOrder string `json:"sortOrder" form:"sortOrder"`
}
