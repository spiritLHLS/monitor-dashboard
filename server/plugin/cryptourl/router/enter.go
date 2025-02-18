package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/cryptourl/api"

var (
	Router           = new(router)
	apiEncryptedLink = api.Api.EncryptedLink
)

type router struct{ EncryptedLink EL }
