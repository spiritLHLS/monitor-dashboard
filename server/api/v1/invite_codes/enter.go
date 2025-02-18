package invite_codes

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ InviteCodesApi }

var invcodeService = service.ServiceGroupApp.Invite_codesServiceGroup.InviteCodesService
