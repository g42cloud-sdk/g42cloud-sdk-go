package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckpointExtraInfoResp struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	RetentionDuration *int32 `json:"retention_duration,omitempty"`
}

func (o CheckpointExtraInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckpointExtraInfoResp struct{}"
	}

	return strings.Join([]string{"CheckpointExtraInfoResp", string(data)}, " ")
}
