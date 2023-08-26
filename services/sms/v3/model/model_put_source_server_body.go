package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PutSourceServerBody struct {
	Name *string `json:"name,omitempty"`

	Migprojectid *string `json:"migprojectid,omitempty"`

	Disks *[]PutDisk `json:"disks,omitempty"`

	VolumeGroups *[]PutVolumeGroups `json:"volume_groups,omitempty"`
}

func (o PutSourceServerBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutSourceServerBody struct{}"
	}

	return strings.Join([]string{"PutSourceServerBody", string(data)}, " ")
}
