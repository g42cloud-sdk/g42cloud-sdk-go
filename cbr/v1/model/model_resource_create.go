package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceCreate struct {
	ExtraInfo *ResourceExtraInfo `json:"extra_info,omitempty"`

	Id string `json:"id"`

	Type string `json:"type"`

	Name *string `json:"name,omitempty"`
}

func (o ResourceCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceCreate struct{}"
	}

	return strings.Join([]string{"ResourceCreate", string(data)}, " ")
}
