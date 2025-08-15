package request

import (
	"server/model/common/request"
	"time"
)

type ProductsSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Tag           string     `json:"tag" form:"tag" `
	Cpu           string     `json:"cpu" form:"cpu" `
	Memory        string     `json:"memory" form:"memory" `
	Disk          string     `json:"disk" form:"disk" `
	Traffic       string     `json:"traffic" form:"traffic" `
	PortSpeed     string     `json:"portSpeed" form:"portSpeed" `
	Location      string     `json:"location" form:"location" `
	Price         string     `json:"price" form:"price" `
	Additional    string     `json:"additional" form:"additional" `
	Url           string     `json:"url" form:"url" `
	BillingType   string     `json:"billingType" form:"billingType" `
	PushStock     *int       `json:"pushStock" form:"pushStock"`
	OldStock      *int       `json:"oldStock" form:"oldStock" `
	Stock         *int       `json:"stock" form:"stock" `
	MultiCheck    *int       `json:"multiCheck" form:"multiCheck" `
	Intervals     *int       `json:"intervals" form:"intervals" `
	MessageId     string     `json:"messageId" form:"messageId" `
	PushIntervals *int       `json:"pushIntervals" form:"pushIntervals" `
	PushTime      *time.Time `json:"pushTime" form:"pushTime" `
	request.PageInfo
}

type DigitalProductsSearch struct {
	Tag string `json:"tag" form:"tag"`

	// CPU 区间查询
	CpuMin *float64 `json:"cpuMin" form:"cpuMin"`
	CpuMax *float64 `json:"cpuMax" form:"cpuMax"`
	Cpu    *float64 `json:"cpu" form:"cpu"` // 保留兼容性

	// 内存区间查询
	MemoryMin *float64 `json:"memoryMin" form:"memoryMin"`
	MemoryMax *float64 `json:"memoryMax" form:"memoryMax"`
	Memory    *float64 `json:"memory" form:"memory"` // 保留兼容性

	// 磁盘区间查询
	DiskMin *float64 `json:"diskMin" form:"diskMin"`
	DiskMax *float64 `json:"diskMax" form:"diskMax"`
	Disk    *float64 `json:"disk" form:"disk"` // 保留兼容性

	// 流量区间查询
	TrafficMin *float64 `json:"trafficMin" form:"trafficMin"`
	TrafficMax *float64 `json:"trafficMax" form:"trafficMax"`
	Traffic    *float64 `json:"traffic" form:"traffic"` // 保留兼容性

	// 端口速度区间查询
	PortSpeedMin *float64 `json:"portSpeedMin" form:"portSpeedMin"`
	PortSpeedMax *float64 `json:"portSpeedMax" form:"portSpeedMax"`
	PortSpeed    *float64 `json:"portSpeed" form:"portSpeed"` // 保留兼容性

	// 价格区间查询
	PriceMin  *float64 `json:"priceMin" form:"priceMin"`
	PriceMax  *float64 `json:"priceMax" form:"priceMax"`
	Price     *float64 `json:"price" form:"price"` // 保留兼容性
	PriceUnit string   `json:"priceUnit" form:"priceUnit"`

	Location   string `json:"location" form:"location"`
	Additional string `json:"additional" form:"additional"`
	Url        string `json:"url" form:"url"`
	Stock      *int   `json:"stock" form:"stock"`

	// 排序字段
	SortBy    string `json:"sortBy" form:"sortBy"`
	SortOrder string `json:"sortOrder" form:"sortOrder"`

	request.PageInfo
}
