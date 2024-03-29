package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type Follow302StatusRequest struct {
	Follow302Status Follow302StatusRequestFollow302Status `json:"follow302_status"`
}

func (o Follow302StatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Follow302StatusRequest struct{}"
	}

	return strings.Join([]string{"Follow302StatusRequest", string(data)}, " ")
}

type Follow302StatusRequestFollow302Status struct {
	value string
}

type Follow302StatusRequestFollow302StatusEnum struct {
	OFF Follow302StatusRequestFollow302Status
	ON  Follow302StatusRequestFollow302Status
}

func GetFollow302StatusRequestFollow302StatusEnum() Follow302StatusRequestFollow302StatusEnum {
	return Follow302StatusRequestFollow302StatusEnum{
		OFF: Follow302StatusRequestFollow302Status{
			value: "off",
		},
		ON: Follow302StatusRequestFollow302Status{
			value: "on",
		},
	}
}

func (c Follow302StatusRequestFollow302Status) Value() string {
	return c.value
}

func (c Follow302StatusRequestFollow302Status) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Follow302StatusRequestFollow302Status) UnmarshalJSON(b []byte) error {
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
