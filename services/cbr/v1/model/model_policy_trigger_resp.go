package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type PolicyTriggerResp struct {
	Id string `json:"id"`

	Name *string `json:"name,omitempty"`

	Properties *PolicyTriggerPropertiesResp `json:"properties"`

	Type *PolicyTriggerRespType `json:"type,omitempty"`
}

func (o PolicyTriggerResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyTriggerResp struct{}"
	}

	return strings.Join([]string{"PolicyTriggerResp", string(data)}, " ")
}

type PolicyTriggerRespType struct {
	value string
}

type PolicyTriggerRespTypeEnum struct {
	TIME PolicyTriggerRespType
}

func GetPolicyTriggerRespTypeEnum() PolicyTriggerRespTypeEnum {
	return PolicyTriggerRespTypeEnum{
		TIME: PolicyTriggerRespType{
			value: "time",
		},
	}
}

func (c PolicyTriggerRespType) Value() string {
	return c.value
}

func (c PolicyTriggerRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyTriggerRespType) UnmarshalJSON(b []byte) error {
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
