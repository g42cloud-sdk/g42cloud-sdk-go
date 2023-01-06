package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CheckpointReplicateParam struct {
	AutoTrigger *bool `json:"auto_trigger,omitempty"`

	DestinationProjectId string `json:"destination_project_id"`

	DestinationRegion string `json:"destination_region"`

	DestinationVaultId string `json:"destination_vault_id"`

	EnableAcceleration *bool `json:"enable_acceleration,omitempty"`

	VaultId string `json:"vault_id"`
}

func (o CheckpointReplicateParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckpointReplicateParam struct{}"
	}

	return strings.Join([]string{"CheckpointReplicateParam", string(data)}, " ")
}
