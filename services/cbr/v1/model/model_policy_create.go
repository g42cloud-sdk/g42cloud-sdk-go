package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type PolicyCreate struct {
	Enabled *bool `json:"enabled,omitempty"`

	Name string `json:"name"`

	OperationDefinition *PolicyoOdCreate `json:"operation_definition"`

	OperationType PolicyCreateOperationType `json:"operation_type"`

	Trigger *PolicyTriggerReq `json:"trigger"`
}

func (o PolicyCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyCreate struct{}"
	}

	return strings.Join([]string{"PolicyCreate", string(data)}, " ")
}

type PolicyCreateOperationType struct {
	value string
}

type PolicyCreateOperationTypeEnum struct {
	BACKUP      PolicyCreateOperationType
	REPLICATION PolicyCreateOperationType
}

func GetPolicyCreateOperationTypeEnum() PolicyCreateOperationTypeEnum {
	return PolicyCreateOperationTypeEnum{
		BACKUP: PolicyCreateOperationType{
			value: "backup",
		},
		REPLICATION: PolicyCreateOperationType{
			value: "replication",
		},
	}
}

func (c PolicyCreateOperationType) Value() string {
	return c.value
}

func (c PolicyCreateOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyCreateOperationType) UnmarshalJSON(b []byte) error {
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
