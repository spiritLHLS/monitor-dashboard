package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// EncryptedLink 加密链接 结构体
type EncryptedLink struct {
    global.GVA_MODEL
    RedirectUrl  string `json:"redirectUrl" form:"redirectUrl" gorm:"column:redirect_url;comment:解密后跳转的目标链接;"`  //跳转链接
    EncryptedUrl  string `json:"encryptedUrl" form:"encryptedUrl" gorm:"column:encrypted_url;comment:加密后的链接;"`  //加密链接
    DecryptionKey  string `json:"decryptionKey" form:"decryptionKey" gorm:"column:decryption_key;comment:用于解密链接的密钥;"`  //解密密钥
}


// TableName 加密链接 EncryptedLink自定义表名 EncryptedLinks
func (EncryptedLink) TableName() string {
    return "EncryptedLinks"
}

