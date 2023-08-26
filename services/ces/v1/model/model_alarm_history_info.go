package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type AlarmHistoryInfo struct {
	AlarmId *string `json:"alarm_id,omitempty"`

	AlarmName *string `json:"alarm_name,omitempty"`

	AlarmDescription *string `json:"alarm_description,omitempty"`

	Metric *MetricInfo `json:"metric,omitempty"`

	Condition *Condition `json:"condition,omitempty"`

	AlarmLevel *int32 `json:"alarm_level,omitempty"`

	AlarmType *string `json:"alarm_type,omitempty"`

	AlarmEnabled *bool `json:"alarm_enabled,omitempty"`

	AlarmActionEnabled *bool `json:"alarm_action_enabled,omitempty"`

	AlarmActions *[]AlarmActions `json:"alarm_actions,omitempty"`

	OkActions *[]AlarmActions `json:"ok_actions,omitempty"`

	InsufficientdataActions *[]AlarmActions `json:"insufficientdata_actions,omitempty"`

	UpdateTime *int64 `json:"update_time,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	TriggerTime *int64 `json:"trigger_time,omitempty"`

	AlarmStatus *string `json:"alarm_status,omitempty"`

	Datapoints *[]DataPointForAlarmHistory `json:"datapoints,omitempty"`

	AdditionalInfo *AdditionalInfo `json:"additional_info,omitempty"`
}

func (o AlarmHistoryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmHistoryInfo struct{}"
	}

	return strings.Join([]string{"AlarmHistoryInfo", string(data)}, " ")
}
