package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type MetricForAlarm struct {
	Namespace string `json:"namespace"`

	MetricName string `json:"metric_name"`

	Dimensions *[]MetricsDimension `json:"dimensions,omitempty"`

	ResourceGroupId *string `json:"resource_group_id,omitempty"`
}

func (o MetricForAlarm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricForAlarm struct{}"
	}

	return strings.Join([]string{"MetricForAlarm", string(data)}, " ")
}
