package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateCheckpointRequest struct {
	Body *VaultBackupReq `json:"body,omitempty"`
}

func (o CreateCheckpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCheckpointRequest struct{}"
	}

	return strings.Join([]string{"CreateCheckpointRequest", string(data)}, " ")
}
