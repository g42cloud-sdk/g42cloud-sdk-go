package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type Filter struct {
	Condition FilterCondition `json:"condition"`

	IsSupportFilter bool `json:"is_support_filter"`

	Rule []string `json:"rule"`
}

func (o Filter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Filter struct{}"
	}

	return strings.Join([]string{"Filter", string(data)}, " ")
}

type FilterCondition struct {
	value string
}

type FilterConditionEnum struct {
	AND FilterCondition
	OR  FilterCondition
}

func GetFilterConditionEnum() FilterConditionEnum {
	return FilterConditionEnum{
		AND: FilterCondition{
			value: "AND",
		},
		OR: FilterCondition{
			value: "OR",
		},
	}
}

func (c FilterCondition) Value() string {
	return c.value
}

func (c FilterCondition) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FilterCondition) UnmarshalJSON(b []byte) error {
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
