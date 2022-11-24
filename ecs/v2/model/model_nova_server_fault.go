package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NovaServerFault struct {
	Code *int32 `json:"code,omitempty"`

	Created *string `json:"created,omitempty"`

	Message *string `json:"message,omitempty"`

	Details *string `json:"details,omitempty"`
}

func (o NovaServerFault) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaServerFault struct{}"
	}

	return strings.Join([]string{"NovaServerFault", string(data)}, " ")
}
