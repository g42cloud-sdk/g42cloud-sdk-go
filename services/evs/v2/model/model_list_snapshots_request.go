package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListSnapshotsRequest struct {
	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *string `json:"status,omitempty"`

	VolumeId *string `json:"volume_id,omitempty"`

	AvailabilityZone *string `json:"availability_zone,omitempty"`

	Id *string `json:"id,omitempty"`

	DedicatedStorageName *string `json:"dedicated_storage_name,omitempty"`

	DedicatedStorageId *string `json:"dedicated_storage_id,omitempty"`

	ServiceType *string `json:"service_type,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListSnapshotsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotsRequest struct{}"
	}

	return strings.Join([]string{"ListSnapshotsRequest", string(data)}, " ")
}
