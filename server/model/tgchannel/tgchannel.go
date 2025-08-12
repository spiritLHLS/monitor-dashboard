// 自动生成模板Tgchannel
package tgchannel

import (
	"server/global"
)

// tgchannel表 结构体  Tgchannel
type Tgchannel struct {
	global.GVA_MODEL
	Channel       string `json:"channel" form:"channel" gorm:"column:channel;comment:;size:500;"`                      //频道地址
	Tgid          string `json:"tgid" form:"tgid" gorm:"column:tgid;comment:;size:191;"`                               //频道ID
	Flag          string `json:"flag" form:"flag" gorm:"column:flag;comment:填ALL WITHOUT就忽略部分不推送，填ONLY就仅推送;size:191;"` //推送类型标签
	Tokens        string `json:"tokens" form:"tokens" gorm:"column:tokens;comment:英文逗号分隔各个token;size:3000;"`           //Tokens
	Tags          string `json:"tags" form:"tags" gorm:"column:tags;comment:;size:3000;type:text;"`                    //推送商家
	AdditionalKey string `json:"additionalKey" form:"additionalKey" gorm:"column:additional_key;comment:;size:200;"`   //其他信息关键词
	CreatedBy     *int   `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:20;"`             //创建者
	UpdatedBy     *int   `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:20;"`             //更新者
	DeletedBy     *int   `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;size:20;"`             //删除者
}

// TableName tgchannel表 Tgchannel自定义表名 tgchannel
func (Tgchannel) TableName() string {
	return "tgchannel"
}
