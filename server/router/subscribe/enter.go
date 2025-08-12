package subscribe

import api "server/api/v1"

type RouterGroup struct{ SubscribeRouter }

var subApi = api.ApiGroupApp.SubscribeApiGroup.SubscribeApi
