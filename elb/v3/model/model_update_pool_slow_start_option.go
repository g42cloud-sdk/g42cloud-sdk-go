package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePoolSlowStartOption struct {
	Enable *bool `json:"enable,omitempty"`

	Duration *int32 `json:"duration,omitempty"`
}

func (o UpdatePoolSlowStartOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePoolSlowStartOption struct{}"
	}

	return strings.Join([]string{"UpdatePoolSlowStartOption", string(data)}, " ")
}
