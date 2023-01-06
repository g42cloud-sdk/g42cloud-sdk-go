package model

import (
	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/sdktime"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"
	"strings"
)

type AlarmHistoryItemV2 struct {
	RecordId *string `json:"record_id,omitempty"`

	AlarmId *string `json:"alarm_id,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *AlarmHistoryItemV2Status `json:"status,omitempty"`

	Level *AlarmHistoryItemV2Level `json:"level,omitempty"`

	Type *AlarmHistoryItemV2Type `json:"type,omitempty"`

	ActionEnabled *bool `json:"action_enabled,omitempty"`

	BeginTime *sdktime.SdkTime `json:"begin_time,omitempty"`

	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	Metric *Metric `json:"metric,omitempty"`

	Condition *AlarmCondition `json:"condition,omitempty"`

	AdditionalInfo *AdditionalInfo `json:"additional_info,omitempty"`

	AlarmActions *[]Notification `json:"alarm_actions,omitempty"`

	OkActions *[]Notification `json:"ok_actions,omitempty"`

	Datapoints *[]DataPointInfo `json:"datapoints,omitempty"`
}

func (o AlarmHistoryItemV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmHistoryItemV2 struct{}"
	}

	return strings.Join([]string{"AlarmHistoryItemV2", string(data)}, " ")
}

type AlarmHistoryItemV2Status struct {
	value string
}

type AlarmHistoryItemV2StatusEnum struct {
	OK      AlarmHistoryItemV2Status
	ALARM   AlarmHistoryItemV2Status
	INVALID AlarmHistoryItemV2Status
}

func GetAlarmHistoryItemV2StatusEnum() AlarmHistoryItemV2StatusEnum {
	return AlarmHistoryItemV2StatusEnum{
		OK: AlarmHistoryItemV2Status{
			value: "ok",
		},
		ALARM: AlarmHistoryItemV2Status{
			value: "alarm",
		},
		INVALID: AlarmHistoryItemV2Status{
			value: "invalid",
		},
	}
}

func (c AlarmHistoryItemV2Status) Value() string {
	return c.value
}

func (c AlarmHistoryItemV2Status) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmHistoryItemV2Status) UnmarshalJSON(b []byte) error {
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

type AlarmHistoryItemV2Level struct {
	value int32
}

type AlarmHistoryItemV2LevelEnum struct {
	E_1 AlarmHistoryItemV2Level
	E_2 AlarmHistoryItemV2Level
	E_3 AlarmHistoryItemV2Level
	E_4 AlarmHistoryItemV2Level
}

func GetAlarmHistoryItemV2LevelEnum() AlarmHistoryItemV2LevelEnum {
	return AlarmHistoryItemV2LevelEnum{
		E_1: AlarmHistoryItemV2Level{
			value: 1,
		}, E_2: AlarmHistoryItemV2Level{
			value: 2,
		}, E_3: AlarmHistoryItemV2Level{
			value: 3,
		}, E_4: AlarmHistoryItemV2Level{
			value: 4,
		},
	}
}

func (c AlarmHistoryItemV2Level) Value() int32 {
	return c.value
}

func (c AlarmHistoryItemV2Level) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmHistoryItemV2Level) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type AlarmHistoryItemV2Type struct {
	value string
}

type AlarmHistoryItemV2TypeEnum struct {
	EVENT_SYS    AlarmHistoryItemV2Type
	EVENT_CUSTOM AlarmHistoryItemV2Type
}

func GetAlarmHistoryItemV2TypeEnum() AlarmHistoryItemV2TypeEnum {
	return AlarmHistoryItemV2TypeEnum{
		EVENT_SYS: AlarmHistoryItemV2Type{
			value: "EVENT.SYS",
		},
		EVENT_CUSTOM: AlarmHistoryItemV2Type{
			value: "EVENT.CUSTOM",
		},
	}
}

func (c AlarmHistoryItemV2Type) Value() string {
	return c.value
}

func (c AlarmHistoryItemV2Type) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmHistoryItemV2Type) UnmarshalJSON(b []byte) error {
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
