package shops

import "server/service"

type ApiGroup struct{ ShopsApi }

var spService = service.ServiceGroupApp.ShopsServiceGroup.ShopsService
