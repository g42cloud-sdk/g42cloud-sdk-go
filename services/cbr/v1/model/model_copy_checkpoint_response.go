package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CopyCheckpointResponse struct {
	Replication    *CheckpointReplicateRespBody `json:"replication,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o CopyCheckpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyCheckpointResponse struct{}"
	}

	return strings.Join([]string{"CopyCheckpointResponse", string(data)}, " ")
}
