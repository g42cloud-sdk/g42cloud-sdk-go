package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Resources struct {
	Quota *int32 `json:"quota,omitempty"`

	Used *int32 `json:"used,omitempty"`

	Type *string `json:"type,omitempty"`
}

func (o Resources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resources struct{}"
	}

	return strings.Join([]string{"Resources", string(data)}, " ")
}
