package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EntrypointItem struct {
	MasterEntrypoint *string `json:"master_entrypoint,omitempty"`

	MasterEntrypointIpv6 *string `json:"master_entrypoint_ipv6,omitempty"`

	SlaveEntrypoint *string `json:"slave_entrypoint,omitempty"`

	SlaveEntrypointIpv6 *string `json:"slave_entrypoint_ipv6,omitempty"`

	Type *EntrypointItemType `json:"type,omitempty"`
}

func (o EntrypointItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EntrypointItem struct{}"
	}

	return strings.Join([]string{"EntrypointItem", string(data)}, " ")
}

type EntrypointItemType struct {
	value string
}

type EntrypointItemTypeEnum struct {
	REGISTRY EntrypointItemType
	SERVICE  EntrypointItemType
}

func GetEntrypointItemTypeEnum() EntrypointItemTypeEnum {
	return EntrypointItemTypeEnum{
		REGISTRY: EntrypointItemType{
			value: "REGISTRY",
		},
		SERVICE: EntrypointItemType{
			value: "SERVICE",
		},
	}
}

func (c EntrypointItemType) Value() string {
	return c.value
}

func (c EntrypointItemType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EntrypointItemType) UnmarshalJSON(b []byte) error {
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
