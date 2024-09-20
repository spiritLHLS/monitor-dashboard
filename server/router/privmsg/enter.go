package privmsg

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ PusherConfigRouter }

var pcApi = api.ApiGroupApp.PrivmsgApiGroup.PusherConfigApi
