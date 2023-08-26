package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type CreateNotificationResponse struct {
	NotificationName *string `json:"notification_name,omitempty"`

	OperationType *CreateNotificationResponseOperationType `json:"operation_type,omitempty"`

	Operations *[]Operations `json:"operations,omitempty"`

	NotifyUserList *[]NotificationUsers `json:"notify_user_list,omitempty"`

	Status *CreateNotificationResponseStatus `json:"status,omitempty"`

	TopicId *string `json:"topic_id,omitempty"`

	NotificationId *string `json:"notification_id,omitempty"`

	NotificationType *CreateNotificationResponseNotificationType `json:"notification_type,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	CreateTime *int64 `json:"create_time,omitempty"`

	Filter         *Filter `json:"filter,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateNotificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNotificationResponse struct{}"
	}

	return strings.Join([]string{"CreateNotificationResponse", string(data)}, " ")
}

type CreateNotificationResponseOperationType struct {
	value string
}

type CreateNotificationResponseOperationTypeEnum struct {
	CUSTOMIZED CreateNotificationResponseOperationType
	COMPLETE   CreateNotificationResponseOperationType
}

func GetCreateNotificationResponseOperationTypeEnum() CreateNotificationResponseOperationTypeEnum {
	return CreateNotificationResponseOperationTypeEnum{
		CUSTOMIZED: CreateNotificationResponseOperationType{
			value: "customized",
		},
		COMPLETE: CreateNotificationResponseOperationType{
			value: "complete",
		},
	}
}

func (c CreateNotificationResponseOperationType) Value() string {
	return c.value
}

func (c CreateNotificationResponseOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateNotificationResponseOperationType) UnmarshalJSON(b []byte) error {
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

type CreateNotificationResponseStatus struct {
	value string
}

type CreateNotificationResponseStatusEnum struct {
	ENABLED  CreateNotificationResponseStatus
	DISABLED CreateNotificationResponseStatus
}

func GetCreateNotificationResponseStatusEnum() CreateNotificationResponseStatusEnum {
	return CreateNotificationResponseStatusEnum{
		ENABLED: CreateNotificationResponseStatus{
			value: "enabled",
		},
		DISABLED: CreateNotificationResponseStatus{
			value: "disabled",
		},
	}
}

func (c CreateNotificationResponseStatus) Value() string {
	return c.value
}

func (c CreateNotificationResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateNotificationResponseStatus) UnmarshalJSON(b []byte) error {
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

type CreateNotificationResponseNotificationType struct {
	value string
}

type CreateNotificationResponseNotificationTypeEnum struct {
	SMN CreateNotificationResponseNotificationType
	FUN CreateNotificationResponseNotificationType
}

func GetCreateNotificationResponseNotificationTypeEnum() CreateNotificationResponseNotificationTypeEnum {
	return CreateNotificationResponseNotificationTypeEnum{
		SMN: CreateNotificationResponseNotificationType{
			value: "smn",
		},
		FUN: CreateNotificationResponseNotificationType{
			value: "fun",
		},
	}
}

func (c CreateNotificationResponseNotificationType) Value() string {
	return c.value
}

func (c CreateNotificationResponseNotificationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateNotificationResponseNotificationType) UnmarshalJSON(b []byte) error {
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
