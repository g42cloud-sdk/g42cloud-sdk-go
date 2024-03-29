package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ListPoliciesRequest struct {
	OperationType *ListPoliciesRequestOperationType `json:"operation_type,omitempty"`

	VaultId *string `json:"vault_id,omitempty"`
}

func (o ListPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoliciesRequest struct{}"
	}

	return strings.Join([]string{"ListPoliciesRequest", string(data)}, " ")
}

type ListPoliciesRequestOperationType struct {
	value string
}

type ListPoliciesRequestOperationTypeEnum struct {
	BACKUP      ListPoliciesRequestOperationType
	REPLICATION ListPoliciesRequestOperationType
}

func GetListPoliciesRequestOperationTypeEnum() ListPoliciesRequestOperationTypeEnum {
	return ListPoliciesRequestOperationTypeEnum{
		BACKUP: ListPoliciesRequestOperationType{
			value: "backup",
		},
		REPLICATION: ListPoliciesRequestOperationType{
			value: "replication",
		},
	}
}

func (c ListPoliciesRequestOperationType) Value() string {
	return c.value
}

func (c ListPoliciesRequestOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPoliciesRequestOperationType) UnmarshalJSON(b []byte) error {
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
