package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type OutputPolicy struct {
	OutputPolicy *OutputPolicyOutputPolicy `json:"output_policy,omitempty"`
}

func (o OutputPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OutputPolicy struct{}"
	}

	return strings.Join([]string{"OutputPolicy", string(data)}, " ")
}

type OutputPolicyOutputPolicy struct {
	value string
}

type OutputPolicyOutputPolicyEnum struct {
	TRANSCODE OutputPolicyOutputPolicy
	DISCARD   OutputPolicyOutputPolicy
	COPY      OutputPolicyOutputPolicy
}

func GetOutputPolicyOutputPolicyEnum() OutputPolicyOutputPolicyEnum {
	return OutputPolicyOutputPolicyEnum{
		TRANSCODE: OutputPolicyOutputPolicy{
			value: "transcode",
		},
		DISCARD: OutputPolicyOutputPolicy{
			value: "discard",
		},
		COPY: OutputPolicyOutputPolicy{
			value: "copy",
		},
	}
}

func (c OutputPolicyOutputPolicy) Value() string {
	return c.value
}

func (c OutputPolicyOutputPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OutputPolicyOutputPolicy) UnmarshalJSON(b []byte) error {
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
