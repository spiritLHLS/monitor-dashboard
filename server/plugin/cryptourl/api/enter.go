package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/cryptourl/service"

var (
	Api                  = new(api)
	serviceEncryptedLink = service.Service.EncryptedLink
)

type api struct{ EncryptedLink EL }
