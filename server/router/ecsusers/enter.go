package ecsusers

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ EcsUsersRouter }

var eusrApi = api.ApiGroupApp.EcsusersApiGroup.EcsUsersApi
