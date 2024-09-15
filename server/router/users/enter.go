package users

import api "github.com/flipped-aurora/gin-vue-admin/api/v1"

type RouterGroup struct{ UserRouter }

var usrApi = api.ApiGroupApp.UsersApiGroup.UserApi
