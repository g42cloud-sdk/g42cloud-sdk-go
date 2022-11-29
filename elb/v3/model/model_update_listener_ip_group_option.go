package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateListenerIpGroupOption struct {
	IpgroupId *string `json:"ipgroup_id,omitempty"`

	EnableIpgroup *bool `json:"enable_ipgroup,omitempty"`

	Type *UpdateListenerIpGroupOptionType `json:"type,omitempty"`
}

func (o UpdateListenerIpGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateListenerIpGroupOption struct{}"
	}

	return strings.Join([]string{"UpdateListenerIpGroupOption", string(data)}, " ")
}

type UpdateListenerIpGroupOptionType struct {
	value string
}

type UpdateListenerIpGroupOptionTypeEnum struct {
	WHITE UpdateListenerIpGroupOptionType
	BLACK UpdateListenerIpGroupOptionType
}

func GetUpdateListenerIpGroupOptionTypeEnum() UpdateListenerIpGroupOptionTypeEnum {
	return UpdateListenerIpGroupOptionTypeEnum{
		WHITE: UpdateListenerIpGroupOptionType{
			value: "white",
		},
		BLACK: UpdateListenerIpGroupOptionType{
			value: "black",
		},
	}
}

func (c UpdateListenerIpGroupOptionType) Value() string {
	return c.value
}

func (c UpdateListenerIpGroupOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateListenerIpGroupOptionType) UnmarshalJSON(b []byte) error {
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
