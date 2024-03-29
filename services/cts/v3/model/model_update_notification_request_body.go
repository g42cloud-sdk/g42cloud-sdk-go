package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type UpdateNotificationRequestBody struct {
	NotificationName string `json:"notification_name"`

	OperationType UpdateNotificationRequestBodyOperationType `json:"operation_type"`

	Operations *[]Operations `json:"operations,omitempty"`

	NotifyUserList *[]NotificationUsers `json:"notify_user_list,omitempty"`

	Status UpdateNotificationRequestBodyStatus `json:"status"`

	TopicId *string `json:"topic_id,omitempty"`

	NotificationId string `json:"notification_id"`

	Filter *Filter `json:"filter,omitempty"`
}

func (o UpdateNotificationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotificationRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateNotificationRequestBody", string(data)}, " ")
}

type UpdateNotificationRequestBodyOperationType struct {
	value string
}

type UpdateNotificationRequestBodyOperationTypeEnum struct {
	CUSTOMIZED UpdateNotificationRequestBodyOperationType
	COMPLETE   UpdateNotificationRequestBodyOperationType
}

func GetUpdateNotificationRequestBodyOperationTypeEnum() UpdateNotificationRequestBodyOperationTypeEnum {
	return UpdateNotificationRequestBodyOperationTypeEnum{
		CUSTOMIZED: UpdateNotificationRequestBodyOperationType{
			value: "customized",
		},
		COMPLETE: UpdateNotificationRequestBodyOperationType{
			value: "complete",
		},
	}
}

func (c UpdateNotificationRequestBodyOperationType) Value() string {
	return c.value
}

func (c UpdateNotificationRequestBodyOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateNotificationRequestBodyOperationType) UnmarshalJSON(b []byte) error {
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

type UpdateNotificationRequestBodyStatus struct {
	value string
}

type UpdateNotificationRequestBodyStatusEnum struct {
	ENABLED  UpdateNotificationRequestBodyStatus
	DISABLED UpdateNotificationRequestBodyStatus
}

func GetUpdateNotificationRequestBodyStatusEnum() UpdateNotificationRequestBodyStatusEnum {
	return UpdateNotificationRequestBodyStatusEnum{
		ENABLED: UpdateNotificationRequestBodyStatus{
			value: "enabled",
		},
		DISABLED: UpdateNotificationRequestBodyStatus{
			value: "disabled",
		},
	}
}

func (c UpdateNotificationRequestBodyStatus) Value() string {
	return c.value
}

func (c UpdateNotificationRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateNotificationRequestBodyStatus) UnmarshalJSON(b []byte) error {
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
