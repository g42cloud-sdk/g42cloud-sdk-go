package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type OriginHostBody struct {
	OriginHostType OriginHostBodyOriginHostType `json:"origin_host_type"`

	CustomizeDomain *string `json:"customize_domain,omitempty"`
}

func (o OriginHostBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OriginHostBody struct{}"
	}

	return strings.Join([]string{"OriginHostBody", string(data)}, " ")
}

type OriginHostBodyOriginHostType struct {
	value string
}

type OriginHostBodyOriginHostTypeEnum struct {
	ACCELERATE OriginHostBodyOriginHostType
	CUSTOMIZE  OriginHostBodyOriginHostType
}

func GetOriginHostBodyOriginHostTypeEnum() OriginHostBodyOriginHostTypeEnum {
	return OriginHostBodyOriginHostTypeEnum{
		ACCELERATE: OriginHostBodyOriginHostType{
			value: "accelerate",
		},
		CUSTOMIZE: OriginHostBodyOriginHostType{
			value: "customize",
		},
	}
}

func (c OriginHostBodyOriginHostType) Value() string {
	return c.value
}

func (c OriginHostBodyOriginHostType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OriginHostBodyOriginHostType) UnmarshalJSON(b []byte) error {
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
