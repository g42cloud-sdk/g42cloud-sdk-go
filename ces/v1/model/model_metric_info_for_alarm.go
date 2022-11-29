package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetricInfoForAlarm struct {
	Namespace string `json:"namespace"`

	MetricName string `json:"metric_name"`

	Dimensions []MetricsDimension `json:"dimensions"`

	ResourceGroupId *string `json:"resource_group_id,omitempty"`

	ResourceGroupName *string `json:"resource_group_name,omitempty"`
}

func (o MetricInfoForAlarm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricInfoForAlarm struct{}"
	}

	return strings.Join([]string{"MetricInfoForAlarm", string(data)}, " ")
}
