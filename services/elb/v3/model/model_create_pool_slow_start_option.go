package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreatePoolSlowStartOption struct {
	Enable *bool `json:"enable,omitempty"`

	Duration *int32 `json:"duration,omitempty"`
}

func (o CreatePoolSlowStartOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePoolSlowStartOption struct{}"
	}

	return strings.Join([]string{"CreatePoolSlowStartOption", string(data)}, " ")
}
