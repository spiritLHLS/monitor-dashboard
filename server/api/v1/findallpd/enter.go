package findallpd

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ FindallpdApi }

var fapdService = service.ServiceGroupApp.FindallpdServiceGroup.FindallpdService
