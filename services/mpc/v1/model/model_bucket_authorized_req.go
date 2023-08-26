package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type BucketAuthorizedReq struct {
	Bucket string `json:"bucket"`

	Operation BucketAuthorizedReqOperation `json:"operation"`
}

func (o BucketAuthorizedReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BucketAuthorizedReq struct{}"
	}

	return strings.Join([]string{"BucketAuthorizedReq", string(data)}, " ")
}

type BucketAuthorizedReqOperation struct {
	value string
}

type BucketAuthorizedReqOperationEnum struct {
	E_0 BucketAuthorizedReqOperation
	E_1 BucketAuthorizedReqOperation
}

func GetBucketAuthorizedReqOperationEnum() BucketAuthorizedReqOperationEnum {
	return BucketAuthorizedReqOperationEnum{
		E_0: BucketAuthorizedReqOperation{
			value: "0",
		},
		E_1: BucketAuthorizedReqOperation{
			value: "1",
		},
	}
}

func (c BucketAuthorizedReqOperation) Value() string {
	return c.value
}

func (c BucketAuthorizedReqOperation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BucketAuthorizedReqOperation) UnmarshalJSON(b []byte) error {
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
