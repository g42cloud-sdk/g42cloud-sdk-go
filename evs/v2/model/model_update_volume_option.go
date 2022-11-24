package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVolumeOption struct {
	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`
}

func (o UpdateVolumeOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVolumeOption struct{}"
	}

	return strings.Join([]string{"UpdateVolumeOption", string(data)}, " ")
}
