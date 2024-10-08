// 自动生成模板PusherConfig
package privmsg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 推送配置 结构体  PusherConfig
type PusherConfig struct {
	global.GVA_MODEL
	PushType              string     `json:"pushType" form:"pushType" gorm:"column:push_type;comment:;"`                                         //推送类型
	ConfigName            string     `json:"configName" form:"configName" gorm:"column:config_name;comment:;size:100;"`                          //配置名称
	ConfigValue           string     `json:"configValue" form:"configValue" gorm:"column:config_value;comment:;size:255;"`                       //配置数值
	MaxCharactersPerPush  *int       `json:"maxCharactersPerPush" form:"maxCharactersPerPush" gorm:"column:max_characters_per_push;comment:;"`   //最大字符数
	MaxPushesPerInterval  *int       `json:"maxPushesPerInterval" form:"maxPushesPerInterval" gorm:"column:max_pushes_per_interval;comment:;"`   //最大推送次数
	IntervalSeconds       *int       `json:"intervalSeconds" form:"intervalSeconds" gorm:"column:interval_seconds;comment:;"`                    //使用间隔
	CurrentIntervalPushes *int       `json:"currentIntervalPushes" form:"currentIntervalPushes" gorm:"column:current_interval_pushes;comment:;"` //已推送次数 暂时没啥用，轮换的时候直接轮换了，用不上
	LastPushTime          *time.Time `json:"lastPushTime" form:"lastPushTime" gorm:"column:last_push_time;comment:;"`                            //最后一次推送时间
	IsActive              *bool      `json:"isActive" form:"isActive" gorm:"column:is_active;comment:;"`                                         //配置激活
}

// TableName 推送配置 PusherConfig自定义表名 pusher_config
func (PusherConfig) TableName() string {
	return "pusher_config"
}
