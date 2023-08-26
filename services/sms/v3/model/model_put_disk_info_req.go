package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PutDiskInfoReq struct {
	Disks *[]ServerDisk `json:"disks,omitempty"`

	Volumegroups *[]VolumeGroups `json:"volumegroups,omitempty"`

	BtrfsList *[]BtrfsFileSystem `json:"btrfs_list,omitempty"`
}

func (o PutDiskInfoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutDiskInfoReq struct{}"
	}

	return strings.Join([]string{"PutDiskInfoReq", string(data)}, " ")
}
