package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ShowAgenciesTaskResponse struct {
	OperateType    *ShowAgenciesTaskResponseOperateType `json:"operate_type,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o ShowAgenciesTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgenciesTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowAgenciesTaskResponse", string(data)}, " ")
}

type ShowAgenciesTaskResponseOperateType struct {
	value string
}

type ShowAgenciesTaskResponseOperateTypeEnum struct {
	CREATED  ShowAgenciesTaskResponseOperateType
	CANCELED ShowAgenciesTaskResponseOperateType
}

func GetShowAgenciesTaskResponseOperateTypeEnum() ShowAgenciesTaskResponseOperateTypeEnum {
	return ShowAgenciesTaskResponseOperateTypeEnum{
		CREATED: ShowAgenciesTaskResponseOperateType{
			value: "CREATED",
		},
		CANCELED: ShowAgenciesTaskResponseOperateType{
			value: "CANCELED",
		},
	}
}

func (c ShowAgenciesTaskResponseOperateType) Value() string {
	return c.value
}

func (c ShowAgenciesTaskResponseOperateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAgenciesTaskResponseOperateType) UnmarshalJSON(b []byte) error {
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