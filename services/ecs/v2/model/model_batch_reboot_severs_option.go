package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type BatchRebootSeversOption struct {
	Servers []ServerId `json:"servers"`

	Type BatchRebootSeversOptionType `json:"type"`
}

func (o BatchRebootSeversOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRebootSeversOption struct{}"
	}

	return strings.Join([]string{"BatchRebootSeversOption", string(data)}, " ")
}

type BatchRebootSeversOptionType struct {
	value string
}

type BatchRebootSeversOptionTypeEnum struct {
	SOFT BatchRebootSeversOptionType
	HARD BatchRebootSeversOptionType
}

func GetBatchRebootSeversOptionTypeEnum() BatchRebootSeversOptionTypeEnum {
	return BatchRebootSeversOptionTypeEnum{
		SOFT: BatchRebootSeversOptionType{
			value: "SOFT",
		},
		HARD: BatchRebootSeversOptionType{
			value: "HARD",
		},
	}
}

func (c BatchRebootSeversOptionType) Value() string {
	return c.value
}

func (c BatchRebootSeversOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchRebootSeversOptionType) UnmarshalJSON(b []byte) error {
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
