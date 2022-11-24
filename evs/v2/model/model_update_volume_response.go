package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVolumeResponse struct {
	Attachments *[]Attachment `json:"attachments,omitempty"`

	AvailabilityZone *string `json:"availability_zone,omitempty"`

	Bootable *string `json:"bootable,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	Id *string `json:"id,omitempty"`

	Links *[]Link `json:"links,omitempty"`

	Metadata *VolumeMetadata `json:"metadata,omitempty"`

	Multiattach *bool `json:"multiattach,omitempty"`

	Name *string `json:"name,omitempty"`

	OsVolHostAttrhost *string `json:"os-vol-host-attr:host,omitempty"`

	OsVolTenantAttrtenantId *string `json:"os-vol-tenant-attr:tenant_id,omitempty"`

	Shareable *string `json:"shareable,omitempty"`

	Size *int32 `json:"size,omitempty"`

	SnapshotId *string `json:"snapshot_id,omitempty"`

	SourceVolid *string `json:"source_volid,omitempty"`

	Status *string `json:"status,omitempty"`

	VolumeImageMetadata *interface{} `json:"volume_image_metadata,omitempty"`

	VolumeType *string `json:"volume_type,omitempty"`

	Description *string `json:"description,omitempty"`

	OsVolumeReplicationextendedStatus *string `json:"os-volume-replication:extended_status,omitempty"`
	HttpStatusCode                    int     `json:"-"`
}

func (o UpdateVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVolumeResponse struct{}"
	}

	return strings.Join([]string{"UpdateVolumeResponse", string(data)}, " ")
}
