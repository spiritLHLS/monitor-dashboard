package findallpd

import "server/service"

type ApiGroup struct{ FindallpdApi }

var fapdService = service.ServiceGroupApp.FindallpdServiceGroup.FindallpdService
