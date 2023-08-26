package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ListProtectableRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Name *string `json:"name,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	ProtectableType ListProtectableRequestProtectableType `json:"protectable_type"`

	Status *string `json:"status,omitempty"`

	Id *string `json:"id,omitempty"`

	ServerId *string `json:"server_id,omitempty"`
}

func (o ListProtectableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectableRequest struct{}"
	}

	return strings.Join([]string{"ListProtectableRequest", string(data)}, " ")
}

type ListProtectableRequestProtectableType struct {
	value string
}

type ListProtectableRequestProtectableTypeEnum struct {
	SERVER ListProtectableRequestProtectableType
	DISK   ListProtectableRequestProtectableType
}

func GetListProtectableRequestProtectableTypeEnum() ListProtectableRequestProtectableTypeEnum {
	return ListProtectableRequestProtectableTypeEnum{
		SERVER: ListProtectableRequestProtectableType{
			value: "server",
		},
		DISK: ListProtectableRequestProtectableType{
			value: "disk",
		},
	}
}

func (c ListProtectableRequestProtectableType) Value() string {
	return c.value
}

func (c ListProtectableRequestProtectableType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProtectableRequestProtectableType) UnmarshalJSON(b []byte) error {
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
