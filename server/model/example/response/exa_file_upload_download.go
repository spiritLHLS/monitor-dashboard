package response

import "server/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
