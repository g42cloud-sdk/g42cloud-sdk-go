package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowCheckpointRequest struct {
	CheckpointId string `json:"checkpoint_id"`
}

func (o ShowCheckpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCheckpointRequest struct{}"
	}

	return strings.Join([]string{"ShowCheckpointRequest", string(data)}, " ")
}
