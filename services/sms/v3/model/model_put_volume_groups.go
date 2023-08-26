package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PutVolumeGroups struct {
	LogicalVolumes *[]PutLogicalVolume `json:"logical_volumes,omitempty"`

	Id string `json:"id"`

	NeedMigration *bool `json:"need_migration,omitempty"`

	AdjustSize *int64 `json:"adjust_size,omitempty"`
}

func (o PutVolumeGroups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutVolumeGroups struct{}"
	}

	return strings.Join([]string{"PutVolumeGroups", string(data)}, " ")
}
