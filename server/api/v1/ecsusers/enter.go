package ecsusers

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ EcsUsersApi }

var eusrService = service.ServiceGroupApp.EcsusersServiceGroup.EcsUsersService
