package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type TargetServerByTask struct {
	BtrfsList *[]BtrfsFileSystem `json:"btrfs_list,omitempty"`

	Disks []TargetDisks `json:"disks"`

	Name string `json:"name"`

	VmId string `json:"vm_id"`

	VolumeGroups *[]VolumeGroups `json:"volume_groups,omitempty"`
}

func (o TargetServerByTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetServerByTask struct{}"
	}

	return strings.Join([]string{"TargetServerByTask", string(data)}, " ")
}
