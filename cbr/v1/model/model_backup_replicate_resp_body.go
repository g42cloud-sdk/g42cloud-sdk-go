package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackupReplicateRespBody struct {
	BackupId *string `json:"backup_id,omitempty"`

	DestinationProjectId *string `json:"destination_project_id,omitempty"`

	DestinationRegion *string `json:"destination_region,omitempty"`

	DestinationVaultId *string `json:"destination_vault_id,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	ProviderId *string `json:"provider_id,omitempty"`

	ReplicationRecordId *string `json:"replication_record_id,omitempty"`

	SourceRegion *string `json:"source_region,omitempty"`
}

func (o BackupReplicateRespBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupReplicateRespBody struct{}"
	}

	return strings.Join([]string{"BackupReplicateRespBody", string(data)}, " ")
}
