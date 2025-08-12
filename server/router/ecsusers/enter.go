package ecsusers

import api "server/api/v1"

type RouterGroup struct{ EcsUsersRouter }

var eusrApi = api.ApiGroupApp.EcsusersApiGroup.EcsUsersApi
