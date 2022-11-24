package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SnapshotDetails struct {
	Id *string `json:"id,omitempty"`

	Status *string `json:"status,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	Metadata *interface{} `json:"metadata,omitempty"`

	VolumeId *string `json:"volume_id,omitempty"`

	Size *int32 `json:"size,omitempty"`

	OsExtendedSnapshotAttributesprojectId *string `json:"os-extended-snapshot-attributes:project_id,omitempty"`

	OsExtendedSnapshotAttributesprogress *string `json:"os-extended-snapshot-attributes:progress,omitempty"`
}

func (o SnapshotDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SnapshotDetails struct{}"
	}

	return strings.Join([]string{"SnapshotDetails", string(data)}, " ")
}
