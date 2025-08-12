package ecsusers

import "server/service"

type ApiGroup struct{ EcsUsersApi }

var eusrService = service.ServiceGroupApp.EcsusersServiceGroup.EcsUsersService
