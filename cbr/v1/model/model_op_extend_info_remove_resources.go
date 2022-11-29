package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpExtendInfoRemoveResources struct {
	FailCount *int32 `json:"fail_count,omitempty"`

	TotalCount *int32 `json:"total_count,omitempty"`

	Resources *[]Resource `json:"resources,omitempty"`
}

func (o OpExtendInfoRemoveResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpExtendInfoRemoveResources struct{}"
	}

	return strings.Join([]string{"OpExtendInfoRemoveResources", string(data)}, " ")
}
