package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ResetPwdRequest struct {
	XLanguage *ResetPwdRequestXLanguage `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *PwdResetRequest `json:"body,omitempty"`
}

func (o ResetPwdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetPwdRequest struct{}"
	}

	return strings.Join([]string{"ResetPwdRequest", string(data)}, " ")
}

type ResetPwdRequestXLanguage struct {
	value string
}

type ResetPwdRequestXLanguageEnum struct {
	ZH_CN ResetPwdRequestXLanguage
	EN_US ResetPwdRequestXLanguage
}

func GetResetPwdRequestXLanguageEnum() ResetPwdRequestXLanguageEnum {
	return ResetPwdRequestXLanguageEnum{
		ZH_CN: ResetPwdRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ResetPwdRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ResetPwdRequestXLanguage) Value() string {
	return c.value
}

func (c ResetPwdRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResetPwdRequestXLanguage) UnmarshalJSON(b []byte) error {
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
