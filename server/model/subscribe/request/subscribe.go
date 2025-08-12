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
	Tag        string `json:"tag" form:"tag" gorm:"column:tag;comment:;size:20;"`                        //tag字段
	Cpu        string `json:"cpu" form:"cpu" gorm:"column:cpu;comment:;size:191;"`                       //cpu字段
	Memory     string `json:"memory" form:"memory" gorm:"column:memory;comment:;size:191;"`              //memory字段
	Disk       string `json:"disk" form:"disk" gorm:"column:disk;comment:;size:191;"`                    //disk字段
	Traffic    string `json:"traffic" form:"traffic" gorm:"column:traffic;comment:;size:191;"`           //traffic字段
	PortSpeed  string `json:"portSpeed" form:"portSpeed" gorm:"column:port_speed;comment:;size:191;"`    //portSpeed字段
	Location   string `json:"location" form:"location" gorm:"column:location;comment:;size:191;"`        //location字段
	Price      string `json:"price" form:"price" gorm:"column:price;comment:;size:191;"`                 //price字段
	Additional string `json:"additional" form:"additional" gorm:"column:additional;comment:;type:text;"` //additional字段
	OldStock   *int   `json:"oldStock" form:"oldStock" gorm:"column:old_stock;comment:;"`                //oldStock字段
	Stock      *int   `json:"stock" form:"stock" gorm:"column:stock;comment:;"`                          //stock字段
}

type SubProductsWithNotify struct {
	global.GVA_MODEL
	Tag           string `json:"tag" form:"tag" gorm:"column:tag;comment:;size:20;"`                              //tag字段
	Cpu           string `json:"cpu" form:"cpu" gorm:"column:cpu;comment:;size:191;"`                             //cpu字段
	Memory        string `json:"memory" form:"memory" gorm:"column:memory;comment:;size:191;"`                    //memory字段
	Disk          string `json:"disk" form:"disk" gorm:"column:disk;comment:;size:191;"`                          //disk字段
	Traffic       string `json:"traffic" form:"traffic" gorm:"column:traffic;comment:;size:191;"`                 //traffic字段
	PortSpeed     string `json:"portSpeed" form:"portSpeed" gorm:"column:port_speed;comment:;size:191;"`          //portSpeed字段
	Location      string `json:"location" form:"location" gorm:"column:location;comment:;size:191;"`              //location字段
	Price         string `json:"price" form:"price" gorm:"column:price;comment:;size:191;"`                       //price字段
	Additional    string `json:"additional" form:"additional" gorm:"column:additional;comment:;type:text;"`       //additional字段
	OldStock      *int   `json:"oldStock" form:"oldStock" gorm:"column:old_stock;comment:;"`                      //oldStock字段
	Stock         *int   `json:"stock" form:"stock" gorm:"column:stock;comment:;"`                                //stock字段
	Status        *int   `json:"status" form:"status" `                                                           // status字段
	NotifyChannel string `json:"notify_channel" form:"notify_channel" gorm:"column:notify_channel;comment:通知渠道;"` //通知渠道
}
