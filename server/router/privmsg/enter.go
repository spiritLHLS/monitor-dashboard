package privmsg

import api "server/api/v1"

type RouterGroup struct{ PusherConfigRouter }

var pcApi = api.ApiGroupApp.PrivmsgApiGroup.PusherConfigApi
