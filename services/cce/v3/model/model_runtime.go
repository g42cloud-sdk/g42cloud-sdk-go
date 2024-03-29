package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type Runtime struct {
	Name *RuntimeName `json:"name,omitempty"`
}

func (o Runtime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Runtime struct{}"
	}

	return strings.Join([]string{"Runtime", string(data)}, " ")
}

type RuntimeName struct {
	value string
}

type RuntimeNameEnum struct {
	DOCKER     RuntimeName
	CONTAINERD RuntimeName
}

func GetRuntimeNameEnum() RuntimeNameEnum {
	return RuntimeNameEnum{
		DOCKER: RuntimeName{
			value: "docker",
		},
		CONTAINERD: RuntimeName{
			value: "containerd",
		},
	}
}

func (c RuntimeName) Value() string {
	return c.value
}

func (c RuntimeName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RuntimeName) UnmarshalJSON(b []byte) error {
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
