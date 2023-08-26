package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type RefreshTaskRequestBody struct {
	Type *RefreshTaskRequestBodyType `json:"type,omitempty"`

	Urls []string `json:"urls"`
}

func (o RefreshTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RefreshTaskRequestBody struct{}"
	}

	return strings.Join([]string{"RefreshTaskRequestBody", string(data)}, " ")
}

type RefreshTaskRequestBodyType struct {
	value string
}

type RefreshTaskRequestBodyTypeEnum struct {
	FILE      RefreshTaskRequestBodyType
	DIRECTORY RefreshTaskRequestBodyType
}

func GetRefreshTaskRequestBodyTypeEnum() RefreshTaskRequestBodyTypeEnum {
	return RefreshTaskRequestBodyTypeEnum{
		FILE: RefreshTaskRequestBodyType{
			value: "file",
		},
		DIRECTORY: RefreshTaskRequestBodyType{
			value: "directory",
		},
	}
}

func (c RefreshTaskRequestBodyType) Value() string {
	return c.value
}

func (c RefreshTaskRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RefreshTaskRequestBodyType) UnmarshalJSON(b []byte) error {
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
