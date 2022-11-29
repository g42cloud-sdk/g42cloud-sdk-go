package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckpointCreateSkippedResource struct {
	Id *string `json:"id,omitempty"`

	Type *string `json:"type,omitempty"`

	Name *string `json:"name,omitempty"`

	Code *string `json:"code,omitempty"`

	Reason *string `json:"reason,omitempty"`
}

func (o CheckpointCreateSkippedResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckpointCreateSkippedResource struct{}"
	}

	return strings.Join([]string{"CheckpointCreateSkippedResource", string(data)}, " ")
}
