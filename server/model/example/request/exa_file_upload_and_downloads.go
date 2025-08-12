package request

import (
	"server/model/common/request"
)

type ExaAttachmentCategorySearch struct {
	ClassId int `json:"classId" form:"classId"`
	request.PageInfo
}
