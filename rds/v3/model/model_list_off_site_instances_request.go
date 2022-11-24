package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListOffSiteInstancesRequest struct {
	ContentType *string `json:"Content-Type,omitempty"`

	XLanguage *ListOffSiteInstancesRequestXLanguage `json:"X-Language,omitempty"`

	Offset *interface{} `json:"offset,omitempty"`

	Limit *interface{} `json:"limit,omitempty"`
}

func (o ListOffSiteInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOffSiteInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListOffSiteInstancesRequest", string(data)}, " ")
}

type ListOffSiteInstancesRequestXLanguage struct {
	value string
}

type ListOffSiteInstancesRequestXLanguageEnum struct {
	ZH_CN ListOffSiteInstancesRequestXLanguage
	EN_US ListOffSiteInstancesRequestXLanguage
}

func GetListOffSiteInstancesRequestXLanguageEnum() ListOffSiteInstancesRequestXLanguageEnum {
	return ListOffSiteInstancesRequestXLanguageEnum{
		ZH_CN: ListOffSiteInstancesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListOffSiteInstancesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListOffSiteInstancesRequestXLanguage) Value() string {
	return c.value
}

func (c ListOffSiteInstancesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListOffSiteInstancesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
