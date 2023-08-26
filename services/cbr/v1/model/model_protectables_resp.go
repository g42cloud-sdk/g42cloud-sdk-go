package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ProtectablesResp struct {
	Children []interface{} `json:"children"`

	Detail *interface{} `json:"detail,omitempty"`

	Id string `json:"id"`

	Name string `json:"name"`

	Protectable *ProtectableResult `json:"protectable"`

	Size *int32 `json:"size,omitempty"`

	Status *ProtectablesRespStatus `json:"status,omitempty"`

	Type string `json:"type"`
}

func (o ProtectablesResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectablesResp struct{}"
	}

	return strings.Join([]string{"ProtectablesResp", string(data)}, " ")
}

type ProtectablesRespStatus struct {
	value string
}

type ProtectablesRespStatusEnum struct {
	ACTIVE  ProtectablesRespStatus
	DELETED ProtectablesRespStatus
	ERROR   ProtectablesRespStatus
}

func GetProtectablesRespStatusEnum() ProtectablesRespStatusEnum {
	return ProtectablesRespStatusEnum{
		ACTIVE: ProtectablesRespStatus{
			value: "active",
		},
		DELETED: ProtectablesRespStatus{
			value: "deleted",
		},
		ERROR: ProtectablesRespStatus{
			value: "error",
		},
	}
}

func (c ProtectablesRespStatus) Value() string {
	return c.value
}

func (c ProtectablesRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProtectablesRespStatus) UnmarshalJSON(b []byte) error {
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
