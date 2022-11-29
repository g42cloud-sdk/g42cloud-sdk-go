package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchMetricData struct {
	Unit *string `json:"unit,omitempty"`

	Datapoints []DatapointForBatchMetric `json:"datapoints"`

	Namespace *string `json:"namespace,omitempty"`

	MetricName string `json:"metric_name"`

	Dimensions *[]MetricsDimension `json:"dimensions,omitempty"`
}

func (o BatchMetricData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchMetricData struct{}"
	}

	return strings.Join([]string{"BatchMetricData", string(data)}, " ")
}
