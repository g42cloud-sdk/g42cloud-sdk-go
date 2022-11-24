package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ExtraDhcpOption struct {
	OptName ExtraDhcpOptionOptName `json:"opt_name"`

	OptValue *string `json:"opt_value,omitempty"`
}

func (o ExtraDhcpOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtraDhcpOption struct{}"
	}

	return strings.Join([]string{"ExtraDhcpOption", string(data)}, " ")
}

type ExtraDhcpOptionOptName struct {
	value string
}

type ExtraDhcpOptionOptNameEnum struct {
	NTP ExtraDhcpOptionOptName
}

func GetExtraDhcpOptionOptNameEnum() ExtraDhcpOptionOptNameEnum {
	return ExtraDhcpOptionOptNameEnum{
		NTP: ExtraDhcpOptionOptName{
			value: "ntp",
		},
	}
}

func (c ExtraDhcpOptionOptName) Value() string {
	return c.value
}

func (c ExtraDhcpOptionOptName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExtraDhcpOptionOptName) UnmarshalJSON(b []byte) error {
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
