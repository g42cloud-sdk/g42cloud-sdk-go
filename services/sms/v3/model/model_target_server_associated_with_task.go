package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type TargetServerAssociatedWithTask struct {
	Id *string `json:"id,omitempty"`

	VmId *string `json:"vm_id,omitempty"`

	Name *string `json:"name,omitempty"`

	Ip *string `json:"ip,omitempty"`

	OsType *TargetServerAssociatedWithTaskOsType `json:"os_type,omitempty"`

	OsVersion *string `json:"os_version,omitempty"`
}

func (o TargetServerAssociatedWithTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetServerAssociatedWithTask struct{}"
	}

	return strings.Join([]string{"TargetServerAssociatedWithTask", string(data)}, " ")
}

type TargetServerAssociatedWithTaskOsType struct {
	value string
}

type TargetServerAssociatedWithTaskOsTypeEnum struct {
	WINDOWS TargetServerAssociatedWithTaskOsType
	LINUX   TargetServerAssociatedWithTaskOsType
}

func GetTargetServerAssociatedWithTaskOsTypeEnum() TargetServerAssociatedWithTaskOsTypeEnum {
	return TargetServerAssociatedWithTaskOsTypeEnum{
		WINDOWS: TargetServerAssociatedWithTaskOsType{
			value: "WINDOWS",
		},
		LINUX: TargetServerAssociatedWithTaskOsType{
			value: "LINUX",
		},
	}
}

func (c TargetServerAssociatedWithTaskOsType) Value() string {
	return c.value
}

func (c TargetServerAssociatedWithTaskOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TargetServerAssociatedWithTaskOsType) UnmarshalJSON(b []byte) error {
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
