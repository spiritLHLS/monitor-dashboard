package invite_codes

import "server/service"

type ApiGroup struct{ InviteCodesApi }

var invcodeService = service.ServiceGroupApp.Invite_codesServiceGroup.InviteCodesService
