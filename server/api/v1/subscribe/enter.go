package subscribe

import "server/service"

type ApiGroup struct{ SubscribeApi }

var subService = service.ServiceGroupApp.SubscribeServiceGroup.SubscribeService
