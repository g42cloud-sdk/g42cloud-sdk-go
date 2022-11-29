package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Resource struct {
	ExtraInfo *ResourceExtraInfo `json:"extra_info,omitempty"`

	Id string `json:"id"`

	Name *string `json:"name,omitempty"`

	Type string `json:"type"`
}

func (o Resource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}
