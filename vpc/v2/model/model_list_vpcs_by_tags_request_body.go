package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListVpcsByTagsRequestBody struct {
	Action ListVpcsByTagsRequestBodyAction `json:"action"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Matches *[]Match `json:"matches,omitempty"`

	Tags *[]ListTag `json:"tags,omitempty"`
}

func (o ListVpcsByTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcsByTagsRequestBody struct{}"
	}

	return strings.Join([]string{"ListVpcsByTagsRequestBody", string(data)}, " ")
}

type ListVpcsByTagsRequestBodyAction struct {
	value string
}

type ListVpcsByTagsRequestBodyActionEnum struct {
	FILTER ListVpcsByTagsRequestBodyAction
	COUNT  ListVpcsByTagsRequestBodyAction
}

func GetListVpcsByTagsRequestBodyActionEnum() ListVpcsByTagsRequestBodyActionEnum {
	return ListVpcsByTagsRequestBodyActionEnum{
		FILTER: ListVpcsByTagsRequestBodyAction{
			value: "filter",
		},
		COUNT: ListVpcsByTagsRequestBodyAction{
			value: "count",
		},
	}
}

func (c ListVpcsByTagsRequestBodyAction) Value() string {
	return c.value
}

func (c ListVpcsByTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVpcsByTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
