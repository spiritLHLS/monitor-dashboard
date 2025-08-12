package privmsg

import "server/service"

type ApiGroup struct{ PusherConfigApi }

var pcService = service.ServiceGroupApp.PrivmsgServiceGroup.PusherConfigService
