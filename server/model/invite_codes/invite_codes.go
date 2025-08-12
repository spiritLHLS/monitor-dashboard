// 自动生成模板InviteCodes
package invite_codes

import (
	"github.com/robfig/cron/v3"
	"server/global"
	"time"
)

// 邀请码 结构体  InviteCodes
type InviteCodes struct {
	global.GVA_MODEL
	Code       string     `json:"code" form:"code" gorm:"column:code;comment:;" binding:"required"` //邀请码
	Status     *int       `json:"status" form:"status" gorm:"column:status;comment:;"`              //使用状态
	Expires_at *time.Time `json:"expires_at" form:"expires_at" gorm:"column:expires_at;comment:;"`  //过期时间
	Used_at    *time.Time `json:"used_at" form:"used_at" gorm:"column:used_at;comment:;"`           //使用时间
	User_uuid  string     `json:"user_uuid" form:"user_uuid" gorm:"column:user_uuid;comment:;"`     //使用者
	Remarks    string     `json:"remarks" form:"remarks" gorm:"column:remarks;comment:;"`           //备注
}

// TableName 邀请码 InviteCodes自定义表名 invite_codes
func (InviteCodes) TableName() string {
	return "invite_codes"
}

type BatchBuildCodes struct {
	Count *int `json:"count"`
}

type BatchExportCodes struct {
	IDs []*int `json:"IDs"`
}

var PublicInviteCodesStatus = true
var TaskId cron.EntryID

type InviteControl struct {
	EnablePublicInvite bool `json:"enable_public_invite"`
}
