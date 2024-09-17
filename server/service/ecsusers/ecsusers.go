package ecsusers

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ecsusers"
	ecsusersReq "github.com/flipped-aurora/gin-vue-admin/server/model/ecsusers/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
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
		//{"nickname", eusr.Nickname, "昵称"},
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
// 参数: info ecsusersReq.EcsUsersSearch - 包含搜索条件和分页信息的结构体
// 返回: list []ecsusers.EcsUsers - 订阅用户列表
//
//	total int64 - 总记录数
//	err error - 错误信息
func (eusrService *EcsUsersService) GetEcsUsersInfoList(info ecsusersReq.EcsUsersSearch) (list []ecsusers.EcsUsers, total int64, err error) {
	// 计算分页参数
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建数据库查询对象
	db := global.GVA_DB.Model(&ecsusers.EcsUsers{})
	// 构建查询条件
	db = eusrService.buildQueryConditions(db, info)
	// 获取总记录数
	if err = db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	// 应用分页
	if limit > 0 {
		db = db.Limit(limit).Offset(offset)
	}
	// 执行查询
	var ecsUsersList []ecsusers.EcsUsers
	err = db.Find(&ecsUsersList).Error
	return ecsUsersList, total, err
}

// buildQueryConditions 构建查询条件
// 参数: db *gorm.DB - 数据库查询对象
//
//	info ecsusersReq.EcsUsersSearch - 搜索条件
//
// 返回: *gorm.DB - 添加了查询条件的数据库查询对象
func (eusrService *EcsUsersService) buildQueryConditions(db *gorm.DB, info ecsusersReq.EcsUsersSearch) *gorm.DB {
	// 时间范围查询
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	// 字符串字段模糊查询
	likeFields := map[string]string{
		"username":       info.Username,
		"nickname":       info.Nickname,
		"password":       info.Password,
		"push_channel1":  info.PushChannel1,
		"push_channel2":  info.PushChannel2,
		"push_channel3":  info.PushChannel3,
		"tg_id":          info.TGID,
		"qq_number":      info.QQNumber,
		"we_chat_number": info.WeChatNumber,
		"email":          info.Email,
		"additional":     info.Additional,
	}
	for field, value := range likeFields {
		if value != "" {
			db = db.Where(field+" LIKE ?", "%"+value+"%")
		}
	}
	// 布尔值精确匹配
	if info.IsFrozen != nil {
		db = db.Where("is_frozen = ?", info.IsFrozen)
	}
	// 等级范围查询
	if info.StartLevel != nil && info.EndLevel != nil {
		db = db.Where("level BETWEEN ? AND ?", info.StartLevel, info.EndLevel)
	}
	return db
}

// AdminChangePassword 方法介绍
// Author [yourname](https://github.com/yourname)
func (eusrService *EcsUsersService) AdminChangePassword(req ecsusersReq.AdminChangePasswordReq) (err error) {
	newPwd := utils.BcryptHash(req.Password)
	// 在对应表中通过user的id查询到后更新用户密码
	db := global.GVA_DB.Model(&ecsusers.EcsUsers{}).
		Where("id = ?", req.UserID).Update("password", newPwd)
	return db.Error
}

// Login 方法介绍
// Author [yourname](https://github.com/yourname)
func (eusrService *EcsUsersService) Login(username, password string) (user ecsusers.EcsUsers, err error) {
	// 用户查找和密码比对
	err = global.GVA_DB.First(&user, "username = ?", username).Error
	if err != nil {
		return user, errors.New("用户不存在")
	}
	ok := utils.BcryptCheck(password, user.Password)
	if !ok {
		return user, errors.New("密码错误")
	}
	return
}

// GetUserInfo 方法介绍
// Author [yourname](https://github.com/yourname)
func (eusrService *EcsUsersService) GetUserInfo(id uint) (user ecsusers.EcsUsers, err error) {
	err = global.GVA_DB.First(&user, id).Error
	return
}

func (eusrService *EcsUsersService) GetEcsUsersPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
