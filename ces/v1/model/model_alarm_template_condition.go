package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmTemplateCondition struct {
	ComparisonOperator string `json:"comparison_operator"`

	Count int32 `json:"count"`

	Filter string `json:"filter"`

	Period int32 `json:"period"`

	Unit *string `json:"unit,omitempty"`

	Value float64 `json:"value"`

	SuppressDuration *int32 `json:"suppress_duration,omitempty"`
}

func (o AlarmTemplateCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmTemplateCondition struct{}"
	}

	return strings.Join([]string{"AlarmTemplateCondition", string(data)}, " ")
}
