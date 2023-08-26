package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type VolumeGroups struct {
	Components *string `json:"components,omitempty"`

	FreeSize *int64 `json:"free_size,omitempty"`

	LogicalVolumes *[]LogicalVolumes `json:"logical_volumes,omitempty"`

	Name *string `json:"name,omitempty"`

	Size *int64 `json:"size,omitempty"`
}

func (o VolumeGroups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeGroups struct{}"
	}

	return strings.Join([]string{"VolumeGroups", string(data)}, " ")
}
