package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BackupRestore struct {
	Mappings *[]BackupRestoreServerMapping `json:"mappings,omitempty"`

	PowerOn *bool `json:"power_on,omitempty"`

	ServerId *string `json:"server_id,omitempty"`

	VolumeId *string `json:"volume_id,omitempty"`

	ResourceId *string `json:"resource_id,omitempty"`
}

func (o BackupRestore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupRestore struct{}"
	}

	return strings.Join([]string{"BackupRestore", string(data)}, " ")
}
