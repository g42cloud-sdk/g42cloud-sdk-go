package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type BatchDeleteSubnetTagsRequestBody struct {
	Action BatchDeleteSubnetTagsRequestBodyAction `json:"action"`

	Tags []ResourceTag `json:"tags"`
}

func (o BatchDeleteSubnetTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSubnetTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteSubnetTagsRequestBody", string(data)}, " ")
}

type BatchDeleteSubnetTagsRequestBodyAction struct {
	value string
}

type BatchDeleteSubnetTagsRequestBodyActionEnum struct {
	DELETE BatchDeleteSubnetTagsRequestBodyAction
}

func GetBatchDeleteSubnetTagsRequestBodyActionEnum() BatchDeleteSubnetTagsRequestBodyActionEnum {
	return BatchDeleteSubnetTagsRequestBodyActionEnum{
		DELETE: BatchDeleteSubnetTagsRequestBodyAction{
			value: "delete",
		},
	}
}

func (c BatchDeleteSubnetTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchDeleteSubnetTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteSubnetTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
