package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateListenerIpGroupOption struct {
	IpgroupId string `json:"ipgroup_id"`

	EnableIpgroup *bool `json:"enable_ipgroup,omitempty"`

	Type *CreateListenerIpGroupOptionType `json:"type,omitempty"`
}

func (o CreateListenerIpGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateListenerIpGroupOption struct{}"
	}

	return strings.Join([]string{"CreateListenerIpGroupOption", string(data)}, " ")
}

type CreateListenerIpGroupOptionType struct {
	value string
}

type CreateListenerIpGroupOptionTypeEnum struct {
	WHITE CreateListenerIpGroupOptionType
	BLACK CreateListenerIpGroupOptionType
}

func GetCreateListenerIpGroupOptionTypeEnum() CreateListenerIpGroupOptionTypeEnum {
	return CreateListenerIpGroupOptionTypeEnum{
		WHITE: CreateListenerIpGroupOptionType{
			value: "white",
		},
		BLACK: CreateListenerIpGroupOptionType{
			value: "black",
		},
	}
}

func (c CreateListenerIpGroupOptionType) Value() string {
	return c.value
}

func (c CreateListenerIpGroupOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateListenerIpGroupOptionType) UnmarshalJSON(b []byte) error {
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
