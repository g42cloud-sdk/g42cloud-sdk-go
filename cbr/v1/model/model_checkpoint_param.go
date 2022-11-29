package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckpointParam struct {
	AutoTrigger *bool `json:"auto_trigger,omitempty"`

	Description *string `json:"description,omitempty"`

	Incremental *bool `json:"incremental,omitempty"`

	Name *string `json:"name,omitempty"`

	Resources *[]string `json:"resources,omitempty"`

	ResourceDetails *[]Resource `json:"resource_details,omitempty"`
}

func (o CheckpointParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckpointParam struct{}"
	}

	return strings.Join([]string{"CheckpointParam", string(data)}, " ")
}
