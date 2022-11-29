package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackupReplicateReqBody struct {
	Description *string `json:"description,omitempty"`

	DestinationProjectId string `json:"destination_project_id"`

	DestinationRegion string `json:"destination_region"`

	DestinationVaultId string `json:"destination_vault_id"`

	EnableAcceleration *bool `json:"enable_acceleration,omitempty"`

	Name *string `json:"name,omitempty"`
}

func (o BackupReplicateReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupReplicateReqBody struct{}"
	}

	return strings.Join([]string{"BackupReplicateReqBody", string(data)}, " ")
}
