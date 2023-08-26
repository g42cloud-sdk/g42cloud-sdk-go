package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type InitTargetServer struct {
	Disks []DiskIntargetServer `json:"disks"`

	VolumeGroups *[]VolumeGroups `json:"volume_groups,omitempty"`
}

func (o InitTargetServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InitTargetServer struct{}"
	}

	return strings.Join([]string{"InitTargetServer", string(data)}, " ")
}
