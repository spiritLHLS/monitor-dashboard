package request

import (
	"server/model/common/request"
	"server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
