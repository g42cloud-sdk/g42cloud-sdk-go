package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type EnvironmentCheck struct {
	Id int64 `json:"id"`

	Params *[]string `json:"params,omitempty"`

	Name string `json:"name"`

	Result EnvironmentCheckResult `json:"result"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorOrWarn *string `json:"error_or_warn,omitempty"`

	ErrorParams *string `json:"error_params,omitempty"`
}

func (o EnvironmentCheck) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvironmentCheck struct{}"
	}

	return strings.Join([]string{"EnvironmentCheck", string(data)}, " ")
}

type EnvironmentCheckResult struct {
	value string
}

type EnvironmentCheckResultEnum struct {
	OK    EnvironmentCheckResult
	WARN  EnvironmentCheckResult
	ERROR EnvironmentCheckResult
}

func GetEnvironmentCheckResultEnum() EnvironmentCheckResultEnum {
	return EnvironmentCheckResultEnum{
		OK: EnvironmentCheckResult{
			value: "OK",
		},
		WARN: EnvironmentCheckResult{
			value: "WARN",
		},
		ERROR: EnvironmentCheckResult{
			value: "ERROR",
		},
	}
}

func (c EnvironmentCheckResult) Value() string {
	return c.value
}

func (c EnvironmentCheckResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EnvironmentCheckResult) UnmarshalJSON(b []byte) error {
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
