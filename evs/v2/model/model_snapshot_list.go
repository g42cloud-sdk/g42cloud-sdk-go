package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SnapshotList struct {
	Id string `json:"id"`

	Status string `json:"status"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	CreatedAt string `json:"created_at"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	Metadata map[string]string `json:"metadata,omitempty"`

	VolumeId string `json:"volume_id"`

	Size int32 `json:"size"`

	OsExtendedSnapshotAttributesprojectId string `json:"os-extended-snapshot-attributes:project_id"`

	OsExtendedSnapshotAttributesprogress string `json:"os-extended-snapshot-attributes:progress"`

	DedicatedStorageId *string `json:"dedicated_storage_id,omitempty"`

	DedicatedStorageName *string `json:"dedicated_storage_name,omitempty"`

	ServiceType string `json:"service_type"`
}

func (o SnapshotList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SnapshotList struct{}"
	}

	return strings.Join([]string{"SnapshotList", string(data)}, " ")
}
