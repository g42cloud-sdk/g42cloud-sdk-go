package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpExtendInfoReplication struct {
	DestinationBackupId *string `json:"destination_backup_id,omitempty"`

	DestinationCheckpointId *string `json:"destination_checkpoint_id,omitempty"`

	DestinationProjectId string `json:"destination_project_id"`

	DestinationRegion string `json:"destination_region"`

	SourceBackupId string `json:"source_backup_id"`

	SourceCheckpointId *string `json:"source_checkpoint_id,omitempty"`

	SourceProjectId string `json:"source_project_id"`

	SourceRegion string `json:"source_region"`

	SourceBackupName *string `json:"source_backup_name,omitempty"`

	DestinationBackupName *string `json:"destination_backup_name,omitempty"`
}

func (o OpExtendInfoReplication) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpExtendInfoReplication struct{}"
	}

	return strings.Join([]string{"OpExtendInfoReplication", string(data)}, " ")
}
