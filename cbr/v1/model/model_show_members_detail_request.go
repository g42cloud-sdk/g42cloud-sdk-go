package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowMembersDetailRequest struct {
	BackupId string `json:"backup_id"`

	DestProjectId *string `json:"dest_project_id,omitempty"`

	ImageId *string `json:"image_id,omitempty"`

	Status *string `json:"status,omitempty"`

	VaultId *string `json:"vault_id,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Sort *string `json:"sort,omitempty"`
}

func (o ShowMembersDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMembersDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowMembersDetailRequest", string(data)}, " ")
}
