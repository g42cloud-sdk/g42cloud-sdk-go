package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CheckpointReplicateRespBody struct {
	Backups []CheckpointReplicateRespbackups `json:"backups"`

	DestinationProjectId string `json:"destination_project_id"`

	DestinationRegion string `json:"destination_region"`

	DestinationVaultId string `json:"destination_vault_id"`

	ProjectId string `json:"project_id"`

	ProviderId string `json:"provider_id"`

	SourceRegion string `json:"source_region"`

	VaultId string `json:"vault_id"`
}

func (o CheckpointReplicateRespBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckpointReplicateRespBody struct{}"
	}

	return strings.Join([]string{"CheckpointReplicateRespBody", string(data)}, " ")
}
