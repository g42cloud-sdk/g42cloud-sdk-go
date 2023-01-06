package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type MetricAlarms struct {
	AlarmName string `json:"alarm_name"`

	AlarmDescription *string `json:"alarm_description,omitempty"`

	Metric *MetricInfoForAlarm `json:"metric"`

	Condition *Condition `json:"condition"`

	AlarmEnabled *bool `json:"alarm_enabled,omitempty"`

	AlarmLevel *int32 `json:"alarm_level,omitempty"`

	AlarmType *MetricAlarmsAlarmType `json:"alarm_type,omitempty"`

	AlarmActionEnabled *bool `json:"alarm_action_enabled,omitempty"`

	AlarmActions *[]AlarmActions `json:"alarm_actions,omitempty"`

	OkActions *[]AlarmActions `json:"ok_actions,omitempty"`

	InsufficientdataActions *[]AlarmActions `json:"insufficientdata_actions,omitempty"`

	AlarmActionBeginTime *string `json:"alarm_action_begin_time,omitempty"`

	AlarmActionEndTime *string `json:"alarm_action_end_time,omitempty"`

	AlarmId string `json:"alarm_id"`

	UpdateTime int64 `json:"update_time"`

	AlarmState string `json:"alarm_state"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o MetricAlarms) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricAlarms struct{}"
	}

	return strings.Join([]string{"MetricAlarms", string(data)}, " ")
}

type MetricAlarmsAlarmType struct {
	value string
}

type MetricAlarmsAlarmTypeEnum struct {
	EVENT_SYS    MetricAlarmsAlarmType
	EVENT_CUSTOM MetricAlarmsAlarmType
}

func GetMetricAlarmsAlarmTypeEnum() MetricAlarmsAlarmTypeEnum {
	return MetricAlarmsAlarmTypeEnum{
		EVENT_SYS: MetricAlarmsAlarmType{
			value: "EVENT.SYS",
		},
		EVENT_CUSTOM: MetricAlarmsAlarmType{
			value: "EVENT.CUSTOM",
		},
	}
}

func (c MetricAlarmsAlarmType) Value() string {
	return c.value
}

func (c MetricAlarmsAlarmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MetricAlarmsAlarmType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
