package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type LogicalVolumes struct {
	BlockCount *int32 `json:"block_count,omitempty"`

	BlockSize *int64 `json:"block_size,omitempty"`

	FileSystem string `json:"file_system"`

	InodeSize int32 `json:"inode_size"`

	DeviceUse *int32 `json:"device_use,omitempty"`

	MountPoint string `json:"mount_point"`

	Name string `json:"name"`

	Size int64 `json:"size"`

	UsedSize int64 `json:"used_size"`

	FreeSize int64 `json:"free_size"`
}

func (o LogicalVolumes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogicalVolumes struct{}"
	}

	return strings.Join([]string{"LogicalVolumes", string(data)}, " ")
}
