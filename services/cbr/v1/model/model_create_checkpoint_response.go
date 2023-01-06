package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateCheckpointResponse struct {
	Checkpoint     *CheckpointCreate `json:"checkpoint,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CreateCheckpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCheckpointResponse struct{}"
	}

	return strings.Join([]string{"CreateCheckpointResponse", string(data)}, " ")
}
