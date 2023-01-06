package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowCheckpointResponse struct {
	Checkpoint     *CheckpointCreate `json:"checkpoint,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowCheckpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCheckpointResponse struct{}"
	}

	return strings.Join([]string{"ShowCheckpointResponse", string(data)}, " ")
}
