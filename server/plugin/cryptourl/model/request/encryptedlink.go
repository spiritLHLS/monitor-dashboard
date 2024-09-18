package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type EncryptedLinkSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    RedirectUrl  string `json:"redirectUrl" form:"redirectUrl" `
    EncryptedUrl  string `json:"encryptedUrl" form:"encryptedUrl" `
    DecryptionKey  string `json:"decryptionKey" form:"decryptionKey" `
    request.PageInfo
}
