package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cryptourl/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/cryptourl/model/request"
)

var EncryptedLink = new(EL)

type EL struct {}
// CreateEncryptedLink 创建加密链接记录
// Author [piexlmax](https://github.com/piexlmax)
func (s *EL) CreateEncryptedLink(EL *model.EncryptedLink) (err error) {
	err = global.GVA_DB.Create(EL).Error
	return err
}

// DeleteEncryptedLink 删除加密链接记录
// Author [piexlmax](https://github.com/piexlmax)
func (s *EL) DeleteEncryptedLink(ID string) (err error) {
	err = global.GVA_DB.Delete(&model.EncryptedLink{},"id = ?",ID).Error
	return err
}

// DeleteEncryptedLinkByIds 批量删除加密链接记录
// Author [piexlmax](https://github.com/piexlmax)
func (s *EL) DeleteEncryptedLinkByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.EncryptedLink{},"id in ?",IDs).Error
	return err
}

// UpdateEncryptedLink 更新加密链接记录
// Author [piexlmax](https://github.com/piexlmax)
func (s *EL) UpdateEncryptedLink(EL model.EncryptedLink) (err error) {
	err = global.GVA_DB.Model(&model.EncryptedLink{}).Where("id = ?",EL.ID).Updates(&EL).Error
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
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.RedirectUrl != "" {
            db = db.Where("redirect_url = ?",info.RedirectUrl)
        }
    if info.EncryptedUrl != "" {
            db = db.Where("encrypted_url = ?",info.EncryptedUrl)
        }
    if info.DecryptionKey != "" {
            db = db.Where("decryption_key = ?",info.DecryptionKey)
        }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	err = db.Find(&ELs).Error
	return  ELs, total, err
}

func (s *EL)GetEncryptedLinkPublic() {

}