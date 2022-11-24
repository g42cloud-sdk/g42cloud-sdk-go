package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VpcRoute struct {
	Id string `json:"id"`

	Destination string `json:"destination"`

	Nexthop string `json:"nexthop"`

	Type VpcRouteType `json:"type"`

	VpcId string `json:"vpc_id"`

	TenantId string `json:"tenant_id"`
}

func (o VpcRoute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcRoute struct{}"
	}

	return strings.Join([]string{"VpcRoute", string(data)}, " ")
}

type VpcRouteType struct {
	value string
}

type VpcRouteTypeEnum struct {
	PEERING VpcRouteType
}

func GetVpcRouteTypeEnum() VpcRouteTypeEnum {
	return VpcRouteTypeEnum{
		PEERING: VpcRouteType{
			value: "peering",
		},
	}
}

func (c VpcRouteType) Value() string {
	return c.value
}

func (c VpcRouteType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VpcRouteType) UnmarshalJSON(b []byte) error {
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
