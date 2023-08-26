package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PhysicalVolume struct {
	DeviceUse *string `json:"device_use,omitempty"`

	FileSystem *string `json:"file_system,omitempty"`

	Index *int32 `json:"index,omitempty"`

	MountPoint *string `json:"mount_point,omitempty"`

	Name *string `json:"name,omitempty"`

	Size *int64 `json:"size,omitempty"`

	UsedSize *int64 `json:"used_size,omitempty"`

	InodeSize *int32 `json:"inode_size,omitempty"`

	Uuid *string `json:"uuid,omitempty"`

	SizePerCluster *int32 `json:"size_per_cluster,omitempty"`
}

func (o PhysicalVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PhysicalVolume struct{}"
	}

	return strings.Join([]string{"PhysicalVolume", string(data)}, " ")
}
