package api

import "server/plugin/cryptourl/service"

var (
	Api                  = new(api)
	serviceEncryptedLink = service.Service.EncryptedLink
)

type api struct{ EncryptedLink EL }
