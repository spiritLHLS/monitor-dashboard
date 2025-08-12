package invite_codes

import api "server/api/v1"

type RouterGroup struct{ InviteCodesRouter }

var invcodeApi = api.ApiGroupApp.Invite_codesApiGroup.InviteCodesApi
