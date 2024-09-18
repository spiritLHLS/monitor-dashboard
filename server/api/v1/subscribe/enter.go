package subscribe

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ SubscribeApi }

var subService = service.ServiceGroupApp.SubscribeServiceGroup.SubscribeService
