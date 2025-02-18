package invite_codes

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ InviteCodesRouter }

var invcodeApi = api.ApiGroupApp.Invite_codesApiGroup.InviteCodesApi
