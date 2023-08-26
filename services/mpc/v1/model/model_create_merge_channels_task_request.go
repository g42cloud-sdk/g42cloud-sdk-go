package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateMergeChannelsTaskRequest struct {
	Body *CreateMergeChannelsReq `json:"body,omitempty"`
}

func (o CreateMergeChannelsTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMergeChannelsTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateMergeChannelsTaskRequest", string(data)}, " ")
}
