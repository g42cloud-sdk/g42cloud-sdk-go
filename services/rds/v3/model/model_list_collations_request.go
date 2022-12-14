package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ListCollationsRequest struct {
	XLanguage *ListCollationsRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListCollationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollationsRequest struct{}"
	}

	return strings.Join([]string{"ListCollationsRequest", string(data)}, " ")
}

type ListCollationsRequestXLanguage struct {
	value string
}

type ListCollationsRequestXLanguageEnum struct {
	ZH_CN ListCollationsRequestXLanguage
	EN_US ListCollationsRequestXLanguage
}

func GetListCollationsRequestXLanguageEnum() ListCollationsRequestXLanguageEnum {
	return ListCollationsRequestXLanguageEnum{
		ZH_CN: ListCollationsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListCollationsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListCollationsRequestXLanguage) Value() string {
	return c.value
}

func (c ListCollationsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCollationsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
