package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type Policy struct {
	Enabled bool `json:"enabled"`

	Id string `json:"id"`

	Name string `json:"name"`

	OperationDefinition *PolicyoOdCreate `json:"operation_definition"`

	OperationType PolicyOperationType `json:"operation_type"`

	Trigger *PolicyTriggerResp `json:"trigger"`

	AssociatedVaults *[]PolicyAssociateVault `json:"associated_vaults,omitempty"`
}

func (o Policy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Policy struct{}"
	}

	return strings.Join([]string{"Policy", string(data)}, " ")
}

type PolicyOperationType struct {
	value string
}

type PolicyOperationTypeEnum struct {
	BACKUP PolicyOperationType
}

func GetPolicyOperationTypeEnum() PolicyOperationTypeEnum {
	return PolicyOperationTypeEnum{
		BACKUP: PolicyOperationType{
			value: "backup",
		},
	}
}

func (c PolicyOperationType) Value() string {
	return c.value
}

func (c PolicyOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyOperationType) UnmarshalJSON(b []byte) error {
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
