package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type TemplateItem struct {
	MetricName string `json:"metric_name"`

	Condition *AlarmTemplateCondition `json:"condition"`

	AlarmLevel *int32 `json:"alarm_level,omitempty"`
}

func (o TemplateItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateItem struct{}"
	}

	return strings.Join([]string{"TemplateItem", string(data)}, " ")
}
