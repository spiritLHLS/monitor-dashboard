package request

import (
	"server/model/common/request"
	"server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
