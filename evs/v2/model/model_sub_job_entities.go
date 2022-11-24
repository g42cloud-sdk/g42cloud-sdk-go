package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubJobEntities struct {
	VolumeType *string `json:"volume_type,omitempty"`

	Size *int32 `json:"size,omitempty"`

	VolumeId *string `json:"volume_id,omitempty"`

	Name *string `json:"name,omitempty"`
}

func (o SubJobEntities) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubJobEntities struct{}"
	}

	return strings.Join([]string{"SubJobEntities", string(data)}, " ")
}
