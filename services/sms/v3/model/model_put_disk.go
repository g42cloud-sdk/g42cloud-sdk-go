package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PutDisk struct {
	NeedMigration *bool `json:"need_migration,omitempty"`

	Id string `json:"id"`

	AdjustSize int64 `json:"adjust_size"`

	PhysicalVolumes *[]PutVolume `json:"physical_volumes,omitempty"`
}

func (o PutDisk) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutDisk struct{}"
	}

	return strings.Join([]string{"PutDisk", string(data)}, " ")
}
