package service

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cryptourl/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cryptourl/model/request"
	"gorm.io/gorm"
)

var EncryptedLink = new(EL)

type EL struct{}

// CreateEncryptedLink 创建加密链接记录
// Author [piexlmax](https://github.com/piexlmax)
func (s *EL) CreateEncryptedLink(EL *model.EncryptedLink) (err error) {
	existingCodes := make(map[string]bool)
	// 首先，获取所有现有的短代码
	db := global.GVA_DB.Model(&model.EncryptedLink{})
	var existingLinks []model.EncryptedLink
	if err := db.Select("short_code").Find(&existingLinks).Error; err != nil {
		return fmt.Errorf("error fetching existing short codes: %w", err)
	}
	for _, link := range existingLinks {
		existingCodes[link.ShortCode] = true
	}
	// 生成唯一的短代码
	for {
		newShortCode := generateShortCode()
		if !existingCodes[newShortCode] {
			EL.ShortCode = newShortCode
			break
		}
	}
	err = global.GVA_DB.Create(EL).Error
	return err
}

// DeleteEncryptedLink 删除加密链接记录
// Author [piexlmax](https://github.com/piexlmax)
func (s *EL) DeleteEncryptedLink(ID string) (err error) {
	err = global.GVA_DB.Delete(&model.EncryptedLink{}, "id = ?", ID).Error
	return err
}

// DeleteEncryptedLinkByIds 批量删除加密链接记录
// Author [piexlmax](https://github.com/piexlmax)
func (s *EL) DeleteEncryptedLinkByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.EncryptedLink{}, "id in ?", IDs).Error
	return err
}

// UpdateEncryptedLink 更新加密链接记录
// Author [piexlmax](https://github.com/piexlmax)
func (s *EL) UpdateEncryptedLink(EL model.EncryptedLink) (err error) {
	err = global.GVA_DB.Model(&model.EncryptedLink{}).Where("id = ?", EL.ID).Updates(&EL).Error
	return err
}

// GetEncryptedLink 根据ID获取加密链接记录
// Author [piexlmax](https://github.com/piexlmax)
func (s *EL) GetEncryptedLink(ID string) (EL model.EncryptedLink, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&EL).Error
	return
}

// GetEncryptedLinkInfoList 分页获取加密链接记录
// Author [piexlmax](https://github.com/piexlmax)
func (s *EL) GetEncryptedLinkInfoList(info request.EncryptedLinkSearch) (list []model.EncryptedLink, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.EncryptedLink{})
	var ELs []model.EncryptedLink
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ProductId != nil {
		db = db.Where("product_id = ?", info.ProductId)
	}
	if info.RedirectUrl != "" {
		db = db.Where("redirect_url LIKE ?", "%"+info.RedirectUrl+"%")
	}
	if info.ShortCode != "" {
		db = db.Where("short_code LIKE ?", "%"+info.ShortCode+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&ELs).Error
	return ELs, total, err
}

// generateShortCode 生成8位的短代码
func generateShortCode() string {
	b := make([]byte, 6)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:8]
}

// generateUniqueShortCodes 生成指定数量的唯一短代码
func generateUniqueShortCodes(db *gorm.DB, count int) ([]string, error) {
	uniqueCodes := make(map[string]bool)
	existingCodes := make(map[string]bool)
	// 首先，获取所有现有的短代码
	var existingLinks []model.EncryptedLink
	if err := db.Select("short_code").Find(&existingLinks).Error; err != nil {
		return nil, fmt.Errorf("error fetching existing short codes: %w", err)
	}
	for _, link := range existingLinks {
		existingCodes[link.ShortCode] = true
	}
	// 生成唯一的短代码
	maxAttempts := count * 10 // 设置最大尝试次数，避免无限循环
	attempts := 0
	for len(uniqueCodes) < count && attempts < maxAttempts {
		code := generateShortCode()
		if !existingCodes[code] && !uniqueCodes[code] {
			uniqueCodes[code] = true
		}
		attempts++
	}
	if len(uniqueCodes) < count {
		return nil, fmt.Errorf("unable to generate %d unique codes after %d attempts", count, maxAttempts)
	}
	// 将 map 转换为 slice
	result := make([]string, 0, count)
	for code := range uniqueCodes {
		result = append(result, code)
	}
	return result, nil
}

// BuildEncryptedUrl 生成加密链接
// Author [yourname](https://github.com/yourname)
func (s *EL) BuildEncryptedUrl(ids request.BuildEncryptedUrlIds) (err error) {
	if len(ids.Ids) == 0 {
		return errors.New("ids is empty")
	}
	db := global.GVA_DB.Model(&model.EncryptedLink{})
	// 获取需要更新的链接
	var links []model.EncryptedLink
	if err := db.Where("id IN ? AND redirect_url != ''", ids.Ids).Find(&links).Error; err != nil {
		return fmt.Errorf("error fetching links: %w", err)
	}
	// 生成足够的唯一短代码
	shortCodes, err := generateUniqueShortCodes(db, len(links))
	if err != nil {
		return fmt.Errorf("failed to generate unique short codes: %w", err)
	}
	// 批量更新链接
	for i := range links {
		links[i].ShortCode = shortCodes[i]
	}
	// 批量更新
	for _, link := range links {
		if err := global.GVA_DB.Model(&model.EncryptedLink{}).Where("id = ?", link.ID).Updates(&link).Error; err != nil {
			return fmt.Errorf("failed to update link for id %d: %w", link.ID, err)
		}
	}
	return nil
}

func (s *EL) GetEncryptedLinkPublic() {

}
