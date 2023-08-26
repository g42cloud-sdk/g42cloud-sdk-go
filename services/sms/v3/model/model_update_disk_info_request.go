package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateDiskInfoRequest struct {
	SourceId string `json:"source_id"`

	Body *PutDiskInfoReq `json:"body,omitempty"`
}

func (o UpdateDiskInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDiskInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdateDiskInfoRequest", string(data)}, " ")
}
