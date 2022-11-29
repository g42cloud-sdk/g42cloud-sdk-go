package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourcesInListResp struct {
	ResourceGroupId *string `json:"resource_group_id,omitempty"`

	ResourceGroupName *string `json:"resource_group_name,omitempty"`

	Dimensions *[]MetricDimension `json:"dimensions,omitempty"`
}

func (o ResourcesInListResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcesInListResp struct{}"
	}

	return strings.Join([]string{"ResourcesInListResp", string(data)}, " ")
}
