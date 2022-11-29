package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Policy struct {
	MetricName string `json:"metric_name"`

	Period int32 `json:"period"`

	Filter string `json:"filter"`

	ComparisonOperator string `json:"comparison_operator"`

	Value float64 `json:"value"`

	Unit *string `json:"unit,omitempty"`

	Count int32 `json:"count"`

	SuppressDuration *int32 `json:"suppress_duration,omitempty"`

	Level *int32 `json:"level,omitempty"`
}

func (o Policy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Policy struct{}"
	}

	return strings.Join([]string{"Policy", string(data)}, " ")
}
