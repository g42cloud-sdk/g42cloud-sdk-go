package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type GetServerRemoteConsoleOption struct {
	Protocol GetServerRemoteConsoleOptionProtocol `json:"protocol"`

	Type GetServerRemoteConsoleOptionType `json:"type"`
}

func (o GetServerRemoteConsoleOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetServerRemoteConsoleOption struct{}"
	}

	return strings.Join([]string{"GetServerRemoteConsoleOption", string(data)}, " ")
}

type GetServerRemoteConsoleOptionProtocol struct {
	value string
}

type GetServerRemoteConsoleOptionProtocolEnum struct {
	VNC GetServerRemoteConsoleOptionProtocol
}

func GetGetServerRemoteConsoleOptionProtocolEnum() GetServerRemoteConsoleOptionProtocolEnum {
	return GetServerRemoteConsoleOptionProtocolEnum{
		VNC: GetServerRemoteConsoleOptionProtocol{
			value: "vnc",
		},
	}
}

func (c GetServerRemoteConsoleOptionProtocol) Value() string {
	return c.value
}

func (c GetServerRemoteConsoleOptionProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetServerRemoteConsoleOptionProtocol) UnmarshalJSON(b []byte) error {
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

type GetServerRemoteConsoleOptionType struct {
	value string
}

type GetServerRemoteConsoleOptionTypeEnum struct {
	NOVNC GetServerRemoteConsoleOptionType
}

func GetGetServerRemoteConsoleOptionTypeEnum() GetServerRemoteConsoleOptionTypeEnum {
	return GetServerRemoteConsoleOptionTypeEnum{
		NOVNC: GetServerRemoteConsoleOptionType{
			value: "novnc",
		},
	}
}

func (c GetServerRemoteConsoleOptionType) Value() string {
	return c.value
}

func (c GetServerRemoteConsoleOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetServerRemoteConsoleOptionType) UnmarshalJSON(b []byte) error {
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
