package shops

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ ShopsApi }

var spService = service.ServiceGroupApp.ShopsServiceGroup.ShopsService
