// 自动生成模板Findallpd
package findallpd
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// findallpd表 结构体  Findallpd
type Findallpd struct {
    global.GVA_MODEL
    Tag  string `json:"tag" form:"tag" gorm:"column:tag;comment:;size:250;"`  //TAG 
    Flag  string `json:"flag" form:"flag" gorm:"column:flag;comment:;size:250;"`  //标志 
    BillingType  string `json:"billingType" form:"billingType" gorm:"column:billing_type;comment:;size:191;"`  //计费类型 
    PdId  *int `json:"pdId" form:"pdId" gorm:"column:pd_id;comment:;size:19;"`  //起始编号 
    EndId  *int `json:"endId" form:"endId" gorm:"column:end_id;comment:;size:19;"`  //结束编号 
    Cpu  string `json:"cpu" form:"cpu" gorm:"column:cpu;comment:;size:191;"`  //CPU 
    Memory  string `json:"memory" form:"memory" gorm:"column:memory;comment:;size:191;"`  //内存 
    Disk  string `json:"disk" form:"disk" gorm:"column:disk;comment:;size:191;"`  //硬盘 
    Traffic  string `json:"traffic" form:"traffic" gorm:"column:traffic;comment:;size:191;"`  //流量 
    PortSpeed  string `json:"portSpeed" form:"portSpeed" gorm:"column:port_speed;comment:;size:191;"`  //带宽 
    Location  string `json:"location" form:"location" gorm:"column:location;comment:;size:191;"`  //地址 
    Price  string `json:"price" form:"price" gorm:"column:price;comment:;size:191;"`  //价格 
    Additional  string `json:"additional" form:"additional" gorm:"column:additional;comment:;size:191;"`  //其他信息 
    Url  string `json:"url" form:"url" gorm:"column:url;comment:;size:191;"`  //链接 
    OldStock  *int `json:"oldStock" form:"oldStock" gorm:"column:old_stock;comment:;size:19;"`  //历史库存 
    Stock  *int `json:"stock" form:"stock" gorm:"column:stock;comment:;size:19;"`  //现有库存 
    MessageId  string `json:"messageId" form:"messageId" gorm:"column:message_id;comment:;size:191;"`  //消息ID 
}


// TableName findallpd表 Findallpd自定义表名 findallpd
func (Findallpd) TableName() string {
    return "findallpd"
}

