
package request

import (
	"server/model/common/request"
	"time"
)

type AgentIpRecordSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    Frozen  *bool `json:"frozen" form:"frozen" `
    request.PageInfo
}
