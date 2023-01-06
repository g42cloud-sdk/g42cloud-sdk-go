package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type UpdateDatabaseRequest struct {
	XLanguage *UpdateDatabaseRequestXLanguage `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *UpdateDatabaseReq `json:"body,omitempty"`
}

func (o UpdateDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatabaseRequest struct{}"
	}

	return strings.Join([]string{"UpdateDatabaseRequest", string(data)}, " ")
}

type UpdateDatabaseRequestXLanguage struct {
	value string
}

type UpdateDatabaseRequestXLanguageEnum struct {
	ZH_CN UpdateDatabaseRequestXLanguage
	EN_US UpdateDatabaseRequestXLanguage
}

func GetUpdateDatabaseRequestXLanguageEnum() UpdateDatabaseRequestXLanguageEnum {
	return UpdateDatabaseRequestXLanguageEnum{
		ZH_CN: UpdateDatabaseRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: UpdateDatabaseRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c UpdateDatabaseRequestXLanguage) Value() string {
	return c.value
}

func (c UpdateDatabaseRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDatabaseRequestXLanguage) UnmarshalJSON(b []byte) error {
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