package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type UpdateAlarmRequestBody struct {
	AlarmName *string `json:"alarm_name,omitempty"`

	AlarmDescription *string `json:"alarm_description,omitempty"`

	Condition *Condition `json:"condition,omitempty"`

	AlarmActionEnabled *bool `json:"alarm_action_enabled,omitempty"`

	AlarmLevel *int32 `json:"alarm_level,omitempty"`

	AlarmType *UpdateAlarmRequestBodyAlarmType `json:"alarm_type,omitempty"`

	AlarmActions *[]AlarmActions `json:"alarm_actions,omitempty"`

	InsufficientdataActions *[]AlarmActions `json:"insufficientdata_actions,omitempty"`

	OkActions *[]AlarmActions `json:"ok_actions,omitempty"`
}

func (o UpdateAlarmRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAlarmRequestBody", string(data)}, " ")
}

type UpdateAlarmRequestBodyAlarmType struct {
	value string
}

type UpdateAlarmRequestBodyAlarmTypeEnum struct {
	EVENT_SYS      UpdateAlarmRequestBodyAlarmType
	EVENT_CUSTOM   UpdateAlarmRequestBodyAlarmType
	RESOURCE_GROUP UpdateAlarmRequestBodyAlarmType
}

func GetUpdateAlarmRequestBodyAlarmTypeEnum() UpdateAlarmRequestBodyAlarmTypeEnum {
	return UpdateAlarmRequestBodyAlarmTypeEnum{
		EVENT_SYS: UpdateAlarmRequestBodyAlarmType{
			value: "EVENT.SYS",
		},
		EVENT_CUSTOM: UpdateAlarmRequestBodyAlarmType{
			value: "EVENT.CUSTOM",
		},
		RESOURCE_GROUP: UpdateAlarmRequestBodyAlarmType{
			value: "RESOURCE_GROUP",
		},
	}
}

func (c UpdateAlarmRequestBodyAlarmType) Value() string {
	return c.value
}

func (c UpdateAlarmRequestBodyAlarmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAlarmRequestBodyAlarmType) UnmarshalJSON(b []byte) error {
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
