package privmsg

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ PusherConfigApi }

var pcService = service.ServiceGroupApp.PrivmsgServiceGroup.PusherConfigService
