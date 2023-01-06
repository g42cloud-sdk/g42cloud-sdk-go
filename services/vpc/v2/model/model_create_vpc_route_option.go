package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type CreateVpcRouteOption struct {
	Destination string `json:"destination"`

	Nexthop string `json:"nexthop"`

	Type CreateVpcRouteOptionType `json:"type"`

	VpcId string `json:"vpc_id"`
}

func (o CreateVpcRouteOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcRouteOption struct{}"
	}

	return strings.Join([]string{"CreateVpcRouteOption", string(data)}, " ")
}

type CreateVpcRouteOptionType struct {
	value string
}

type CreateVpcRouteOptionTypeEnum struct {
	PEERING CreateVpcRouteOptionType
}

func GetCreateVpcRouteOptionTypeEnum() CreateVpcRouteOptionTypeEnum {
	return CreateVpcRouteOptionTypeEnum{
		PEERING: CreateVpcRouteOptionType{
			value: "peering",
		},
	}
}

func (c CreateVpcRouteOptionType) Value() string {
	return c.value
}

func (c CreateVpcRouteOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateVpcRouteOptionType) UnmarshalJSON(b []byte) error {
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
