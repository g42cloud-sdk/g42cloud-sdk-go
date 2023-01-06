package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateSnapshotOption struct {
	VolumeId string `json:"volume_id"`

	Force *bool `json:"force,omitempty"`

	Metadata map[string]string `json:"metadata,omitempty"`

	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`
}

func (o CreateSnapshotOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSnapshotOption struct{}"
	}

	return strings.Join([]string{"CreateSnapshotOption", string(data)}, " ")
}
