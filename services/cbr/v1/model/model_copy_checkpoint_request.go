package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CopyCheckpointRequest struct {
	Body *CheckpointReplicateReq `json:"body,omitempty"`
}

func (o CopyCheckpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyCheckpointRequest struct{}"
	}

	return strings.Join([]string{"CopyCheckpointRequest", string(data)}, " ")
}
