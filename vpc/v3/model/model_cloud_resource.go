package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CloudResource struct {
	ResourceType string `json:"resource_type"`

	ResourceCount int32 `json:"resource_count"`
}

func (o CloudResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudResource struct{}"
	}

	return strings.Join([]string{"CloudResource", string(data)}, " ")
}
