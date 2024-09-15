package ecsusers

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ecsusers"
	ecsusersReq "github.com/flipped-aurora/gin-vue-admin/server/model/ecsusers/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gofrs/uuid/v5"
)

type EcsUsersService struct{}

// CreateEcsUsers 创建订阅用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (eusrService *EcsUsersService) CreateEcsUsers(eusr *ecsusers.EcsUsers) error {
	// 验证必填字段
	if err := validateRequiredFields(eusr); err != nil {
		return err
	}
	// 检查唯一性约束
	if err := checkUniqueness(eusr); err != nil {
		return err
	}
	// 密码加密
	eusr.Password = utils.BcryptHash(eusr.Password)
	// 生成 UUID
	var err error
	eusr.UUID, err = uuid.NewV4()
	if err != nil {
		return fmt.Errorf("生成 UUID 失败: %w", err)
	}
	// 创建用户
	return global.GVA_DB.Create(eusr).Error
}

// 验证必填字段
func validateRequiredFields(eusr *ecsusers.EcsUsers) error {
	if eusr.Username == "" {
		return errors.New("用户名不能为空")
	}
	if eusr.Nickname == "" {
		return errors.New("昵称不能为空")
	}
	if eusr.TGID == "" {
		return errors.New("TGID 不能为空")
	}
	return nil
}

// 检查唯一性约束
func checkUniqueness(eusr *ecsusers.EcsUsers) error {
	// 定义需要检查唯一性的字段
	checks := []struct {
		field string
		value string
		label string
	}{
		{"username", eusr.Username, "用户名"},
		{"tg_id", eusr.TGID, "TGID"},
		{"nickname", eusr.Nickname, "昵称"},
	}
	// 如果邮箱不为空，也需要检查唯一性
	if eusr.Email != "" {
		checks = append(checks, struct{ field, value, label string }{"email", eusr.Email, "邮箱"})
	}
	// 遍历检查每个字段的唯一性
	for _, check := range checks {
		var count int64
		err := global.GVA_DB.Model(&ecsusers.EcsUsers{}).Where(check.field+" = ?", check.value).Count(&count).Error
		if err != nil {
			return fmt.Errorf("检查%s时数据库错误: %w", check.label, err)
		}
		if count > 0 {
			return fmt.Errorf("%s (%s) 已存在", check.label, check.value)
		}
	}
	return nil
}

// DeleteEcsUsers 删除订阅用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (eusrService *EcsUsersService) DeleteEcsUsers(ID string) (err error) {
	err = global.GVA_DB.Delete(&ecsusers.EcsUsers{}, "id = ?", ID).Error
	return err
}

// DeleteEcsUsersByIds 批量删除订阅用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (eusrService *EcsUsersService) DeleteEcsUsersByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]ecsusers.EcsUsers{}, "id in ?", IDs).Error
	return err
}

// UpdateEcsUsers 更新订阅用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (eusrService *EcsUsersService) UpdateEcsUsers(eusr ecsusers.EcsUsers) (err error) {
	err = global.GVA_DB.Model(&ecsusers.EcsUsers{}).Where("id = ?", eusr.ID).Updates(&eusr).Error
	return err
}

// GetEcsUsers 根据ID获取订阅用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (eusrService *EcsUsersService) GetEcsUsers(ID string) (eusr ecsusers.EcsUsers, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&eusr).Error
	return
}

// GetEcsUsersInfoList 分页获取订阅用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (eusrService *EcsUsersService) GetEcsUsersInfoList(info ecsusersReq.EcsUsersSearch) (list []ecsusers.EcsUsers, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ecsusers.EcsUsers{})
	var eusrs []ecsusers.EcsUsers
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Username != "" {
		db = db.Where("username LIKE ?", "%"+info.Username+"%")
	}
	if info.Nickname != "" {
		db = db.Where("nickname LIKE ?", "%"+info.Nickname+"%")
	}
	if info.Password != "" {
		db = db.Where("password LIKE ?", "%"+info.Password+"%")
	}
	if info.IsFrozen != nil {
		db = db.Where("is_frozen = ?", info.IsFrozen)
	}
	if info.PushChannel1 != "" {
		db = db.Where("push_channel1 LIKE ?", "%"+info.PushChannel1+"%")
	}
	if info.PushChannel2 != "" {
		db = db.Where("push_channel2 LIKE ?", "%"+info.PushChannel2+"%")
	}
	if info.PushChannel3 != "" {
		db = db.Where("push_channel3 LIKE ?", "%"+info.PushChannel3+"%")
	}
	if info.TGID != "" {
		db = db.Where("tg_id LIKE ?", "%"+info.TGID+"%")
	}
	if info.QQNumber != "" {
		db = db.Where("qq_number LIKE ?", "%"+info.QQNumber+"%")
	}
	if info.WeChatNumber != "" {
		db = db.Where("we_chat_number LIKE ?", "%"+info.WeChatNumber+"%")
	}
	if info.Email != "" {
		db = db.Where("email LIKE ?", "%"+info.Email+"%")
	}
	if info.Additional != "" {
		db = db.Where("additional LIKE ?", "%"+info.Additional+"%")
	}
	if info.StartLevel != nil && info.EndLevel != nil {
		db = db.Where("level BETWEEN ? AND ? ", info.StartLevel, info.EndLevel)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&eusrs).Error
	return eusrs, total, err
}

func (eusrService *EcsUsersService) GetEcsUsersPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
