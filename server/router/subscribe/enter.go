package subscribe

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ SubscribeRouter }

var subApi = api.ApiGroupApp.SubscribeApiGroup.SubscribeApi
