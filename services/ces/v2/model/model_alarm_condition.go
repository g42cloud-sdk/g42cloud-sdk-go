package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type AlarmCondition struct {
	Period AlarmConditionPeriod `json:"period"`

	Filter string `json:"filter"`

	ComparisonOperator string `json:"comparison_operator"`

	Value float64 `json:"value"`

	Unit *string `json:"unit,omitempty"`

	Count int32 `json:"count"`

	SuppressDuration *AlarmConditionSuppressDuration `json:"suppress_duration,omitempty"`
}

func (o AlarmCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmCondition struct{}"
	}

	return strings.Join([]string{"AlarmCondition", string(data)}, " ")
}

type AlarmConditionPeriod struct {
	value int32
}

type AlarmConditionPeriodEnum struct {
	E_0     AlarmConditionPeriod
	E_1     AlarmConditionPeriod
	E_300   AlarmConditionPeriod
	E_1200  AlarmConditionPeriod
	E_3600  AlarmConditionPeriod
	E_14400 AlarmConditionPeriod
	E_86400 AlarmConditionPeriod
}

func GetAlarmConditionPeriodEnum() AlarmConditionPeriodEnum {
	return AlarmConditionPeriodEnum{
		E_0: AlarmConditionPeriod{
			value: 0,
		}, E_1: AlarmConditionPeriod{
			value: 1,
		}, E_300: AlarmConditionPeriod{
			value: 300,
		}, E_1200: AlarmConditionPeriod{
			value: 1200,
		}, E_3600: AlarmConditionPeriod{
			value: 3600,
		}, E_14400: AlarmConditionPeriod{
			value: 14400,
		}, E_86400: AlarmConditionPeriod{
			value: 86400,
		},
	}
}

func (c AlarmConditionPeriod) Value() int32 {
	return c.value
}

func (c AlarmConditionPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmConditionPeriod) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type AlarmConditionSuppressDuration struct {
	value int32
}

type AlarmConditionSuppressDurationEnum struct {
	E_0     AlarmConditionSuppressDuration
	E_300   AlarmConditionSuppressDuration
	E_600   AlarmConditionSuppressDuration
	E_900   AlarmConditionSuppressDuration
	E_1800  AlarmConditionSuppressDuration
	E_3600  AlarmConditionSuppressDuration
	E_10800 AlarmConditionSuppressDuration
	E_21600 AlarmConditionSuppressDuration
	E_43200 AlarmConditionSuppressDuration
}

func GetAlarmConditionSuppressDurationEnum() AlarmConditionSuppressDurationEnum {
	return AlarmConditionSuppressDurationEnum{
		E_0: AlarmConditionSuppressDuration{
			value: 0,
		}, E_300: AlarmConditionSuppressDuration{
			value: 300,
		}, E_600: AlarmConditionSuppressDuration{
			value: 600,
		}, E_900: AlarmConditionSuppressDuration{
			value: 900,
		}, E_1800: AlarmConditionSuppressDuration{
			value: 1800,
		}, E_3600: AlarmConditionSuppressDuration{
			value: 3600,
		}, E_10800: AlarmConditionSuppressDuration{
			value: 10800,
		}, E_21600: AlarmConditionSuppressDuration{
			value: 21600,
		}, E_43200: AlarmConditionSuppressDuration{
			value: 43200,
		},
	}
}

func (c AlarmConditionSuppressDuration) Value() int32 {
	return c.value
}

func (c AlarmConditionSuppressDuration) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmConditionSuppressDuration) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
