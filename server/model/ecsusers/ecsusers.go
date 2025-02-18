// 自动生成模板EcsUsers
package ecsusers

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/google/uuid"
)

// 订阅用户 结构体  EcsUsers
type EcsUsers struct {
	global.GVA_MODEL
	UUID         uuid.UUID `json:"uuid" form:"uuid" gorm:"column:uuid;comment:用户唯一标识;"`                                                                                 //唯一标识
	Username     string    `json:"username" form:"username" gorm:"column:username;comment:用户登录名;" binding:"required"`                                                   //用户名
	Nickname     string    `json:"nickname" form:"nickname" gorm:"default:新用户;column:nickname;comment:用户昵称;" binding:"required"`                                        //昵称
	Avatar       string    `json:"avatar" form:"avatar" gorm:"default:https://raw.githubusercontent.com/spiritlhls/pages/main/logo.png;column:avatar;comment:用户头像URL;"` //头像
	Password     string    `json:"password" form:"password" gorm:"column:password;comment:用户密码;" binding:"required"`                                                    //密码
	IsFrozen     *bool     `json:"isFrozen" form:"isFrozen" gorm:"column:is_frozen;comment:用户是否被冻结;"`                                                                   //是否冻结
	PushChannel1 string    `json:"pushChannel1" form:"pushChannel1" gorm:"column:push_channel1;comment:推送渠道1;"`                                                         //推送渠道1
	PushChannel2 string    `json:"pushChannel2" form:"pushChannel2" gorm:"column:push_channel2;comment:推送渠道2;"`                                                         //推送渠道2
	PushChannel3 string    `json:"pushChannel3" form:"pushChannel3" gorm:"column:push_channel3;comment:推送渠道3;"`                                                         //推送渠道3
	TGID         string    `json:"tgID" form:"tgID" gorm:"column:tg_id;comment:Telegram ID;" binding:"required"`                                                        //TGID
	QQNumber     string    `json:"qqNumber" form:"qqNumber" gorm:"column:qq_number;comment:QQ号码;"`                                                                      //QQ号
	WeChatNumber string    `json:"weChatNumber" form:"weChatNumber" gorm:"column:we_chat_number;comment:微信号码;"`                                                         //微信号
	Email        string    `json:"email" form:"email" gorm:"column:email;comment:用户邮箱;"`                                                                                //邮箱
	Additional   string    `json:"additional" form:"additional" gorm:"column:additional;comment:备注;size:255;"`                                                          //备注
	Level        *int      `json:"level" form:"level" gorm:"default:0;column:level;comment:用户等级;"`                                                                      //用户等级
	AuthorityID  uint      `json:"authorityID" form:"authorityID" gorm:"default:8881;column:authority_id;comment:用户角色;"`                                                //用户角色
}

// 订阅用户客户端可自行更新的内容
type EcsSubUsers struct {
	UUID     uuid.UUID `json:"uuid" form:"uuid" gorm:"column:uuid;comment:用户唯一标识;" binding:"required"`                                                              //唯一标识
	Nickname string    `json:"nickname" form:"nickname" gorm:"default:新用户;column:nickname;comment:用户昵称;" binding:"required"`                                        //昵称
	Avatar   string    `json:"avatar" form:"avatar" gorm:"default:https://raw.githubusercontent.com/spiritlhls/pages/main/logo.png;column:avatar;comment:用户头像URL;"` //头像
	TGID     string    `json:"tgID" form:"tgID" gorm:"column:tg_id;comment:Telegram ID;"`                                                                           //TGID
	Email    string    `json:"email" form:"email" gorm:"column:email;comment:用户邮箱;"`                                                                                //邮箱
	Password string    `json:"password" form:"password" gorm:"column:password;comment:用户密码;"`                                                                       //密码
	//QQNumber     string    `json:"qqNumber" form:"qqNumber" gorm:"column:qq_number;comment:QQ号码;"`                                                                      //QQ号
	//WeChatNumber string    `json:"weChatNumber" form:"weChatNumber" gorm:"column:we_chat_number;comment:微信号码;"`                                                         //微信号
	//PushChannel1 string    `json:"pushChannel1" form:"pushChannel1" gorm:"column:push_channel1;comment:推送渠道1;"`                                                         //推送渠道1
	//PushChannel2 string    `json:"pushChannel2" form:"pushChannel2" gorm:"column:push_channel2;comment:推送渠道2;"`                                                         //推送渠道2
	//PushChannel3 string    `json:"pushChannel3" form:"pushChannel3" gorm:"column:push_channel3;comment:推送渠道3;"`                                                         //推送渠道3
}

func (u EcsUsers) GetUsername() string {
	return u.Username
}

func (u EcsUsers) GetNickname() string {
	return u.Nickname
}

func (u EcsUsers) GetUUID() uuid.UUID {
	return u.UUID
}

func (u EcsUsers) GetUserId() uint {
	return u.ID
}

func (u EcsUsers) GetAuthorityId() uint {
	return u.AuthorityID
}

func (u EcsUsers) GetUserInfo() any {
	return u
}

// TableName 订阅用户 EcsUsers自定义表名 ecs_users
func (EcsUsers) TableName() string {
	return "ecs_users"
}

var EcsUserPublicRegisterStatus = false

type RegisterControl struct {
	EnablePublicRegister bool `json:"enable_public_register"`
}
