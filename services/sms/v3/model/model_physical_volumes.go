package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PhysicalVolumes struct {
	DeviceUse *string `json:"device_use,omitempty"`

	FileSystem *string `json:"file_system,omitempty"`

	Index *int32 `json:"index,omitempty"`

	MountPoint *string `json:"mount_point,omitempty"`

	Name *string `json:"name,omitempty"`

	Size *int64 `json:"size,omitempty"`

	InodeSize *int64 `json:"inode_size,omitempty"`

	UsedSize *int64 `json:"used_size,omitempty"`

	Uuid *string `json:"uuid,omitempty"`
}

func (o PhysicalVolumes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PhysicalVolumes struct{}"
	}

	return strings.Join([]string{"PhysicalVolumes", string(data)}, " ")
}
