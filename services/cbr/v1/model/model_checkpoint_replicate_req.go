package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CheckpointReplicateReq struct {
	Replicate *CheckpointReplicateParam `json:"replicate"`
}

func (o CheckpointReplicateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckpointReplicateReq struct{}"
	}

	return strings.Join([]string{"CheckpointReplicateReq", string(data)}, " ")
}
