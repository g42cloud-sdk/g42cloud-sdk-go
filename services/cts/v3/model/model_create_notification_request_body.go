package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type CreateNotificationRequestBody struct {
	NotificationName string `json:"notification_name"`

	OperationType CreateNotificationRequestBodyOperationType `json:"operation_type"`

	Operations *[]Operations `json:"operations,omitempty"`

	NotifyUserList *[]NotificationUsers `json:"notify_user_list,omitempty"`

	TopicId *string `json:"topic_id,omitempty"`

	Filter *Filter `json:"filter,omitempty"`
}

func (o CreateNotificationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNotificationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateNotificationRequestBody", string(data)}, " ")
}

type CreateNotificationRequestBodyOperationType struct {
	value string
}

type CreateNotificationRequestBodyOperationTypeEnum struct {
	COMPLETE   CreateNotificationRequestBodyOperationType
	CUSTOMIZED CreateNotificationRequestBodyOperationType
}

func GetCreateNotificationRequestBodyOperationTypeEnum() CreateNotificationRequestBodyOperationTypeEnum {
	return CreateNotificationRequestBodyOperationTypeEnum{
		COMPLETE: CreateNotificationRequestBodyOperationType{
			value: "complete",
		},
		CUSTOMIZED: CreateNotificationRequestBodyOperationType{
			value: "customized",
		},
	}
}

func (c CreateNotificationRequestBodyOperationType) Value() string {
	return c.value
}

func (c CreateNotificationRequestBodyOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateNotificationRequestBodyOperationType) UnmarshalJSON(b []byte) error {
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
