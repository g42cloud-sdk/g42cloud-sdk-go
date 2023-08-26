package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateCopyStateRequest struct {
	SourceId string `json:"source_id"`

	Body *PutCopyStateReq `json:"body,omitempty"`
}

func (o UpdateCopyStateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCopyStateRequest struct{}"
	}

	return strings.Join([]string{"UpdateCopyStateRequest", string(data)}, " ")
}
