package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

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