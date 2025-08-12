package router

import "server/plugin/cryptourl/api"

var (
	Router           = new(router)
	apiEncryptedLink = api.Api.EncryptedLink
)

type router struct{ EncryptedLink EL }
