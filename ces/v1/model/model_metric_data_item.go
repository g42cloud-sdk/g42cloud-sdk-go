package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetricDataItem struct {
	Metric *MetricInfo `json:"metric"`

	Ttl int32 `json:"ttl"`

	CollectTime int64 `json:"collect_time"`

	Value float64 `json:"value"`

	Unit *string `json:"unit,omitempty"`

	Type *string `json:"type,omitempty"`
}

func (o MetricDataItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricDataItem struct{}"
	}

	return strings.Join([]string{"MetricDataItem", string(data)}, " ")
}
