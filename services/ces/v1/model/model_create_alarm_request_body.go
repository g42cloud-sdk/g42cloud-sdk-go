package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type CreateAlarmRequestBody struct {
	AlarmName string `json:"alarm_name"`

	AlarmDescription *string `json:"alarm_description,omitempty"`

	Metric *MetricForAlarm `json:"metric"`

	Condition *Condition `json:"condition"`

	AlarmEnabled *bool `json:"alarm_enabled,omitempty"`

	AlarmActionEnabled *bool `json:"alarm_action_enabled,omitempty"`

	AlarmLevel *int32 `json:"alarm_level,omitempty"`

	AlarmType *CreateAlarmRequestBodyAlarmType `json:"alarm_type,omitempty"`

	AlarmActions *[]AlarmActions `json:"alarm_actions,omitempty"`

	InsufficientdataActions *[]AlarmActions `json:"insufficientdata_actions,omitempty"`

	OkActions *[]AlarmActions `json:"ok_actions,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateAlarmRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlarmRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAlarmRequestBody", string(data)}, " ")
}

type CreateAlarmRequestBodyAlarmType struct {
	value string
}

type CreateAlarmRequestBodyAlarmTypeEnum struct {
	EVENT_SYS      CreateAlarmRequestBodyAlarmType
	EVENT_CUSTOM   CreateAlarmRequestBodyAlarmType
	RESOURCE_GROUP CreateAlarmRequestBodyAlarmType
}

func GetCreateAlarmRequestBodyAlarmTypeEnum() CreateAlarmRequestBodyAlarmTypeEnum {
	return CreateAlarmRequestBodyAlarmTypeEnum{
		EVENT_SYS: CreateAlarmRequestBodyAlarmType{
			value: "EVENT.SYS",
		},
		EVENT_CUSTOM: CreateAlarmRequestBodyAlarmType{
			value: "EVENT.CUSTOM",
		},
		RESOURCE_GROUP: CreateAlarmRequestBodyAlarmType{
			value: "RESOURCE_GROUP",
		},
	}
}

func (c CreateAlarmRequestBodyAlarmType) Value() string {
	return c.value
}

func (c CreateAlarmRequestBodyAlarmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAlarmRequestBodyAlarmType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
