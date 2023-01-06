package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type NotificationsResponseBody struct {
	NotificationName *string `json:"notification_name,omitempty"`

	OperationType *NotificationsResponseBodyOperationType `json:"operation_type,omitempty"`

	Operations *[]Operations `json:"operations,omitempty"`

	NotifyUserList *[]NotificationUsers `json:"notify_user_list,omitempty"`

	Status *NotificationsResponseBodyStatus `json:"status,omitempty"`

	TopicId *string `json:"topic_id,omitempty"`

	NotificationId *string `json:"notification_id,omitempty"`

	NotificationType *NotificationsResponseBodyNotificationType `json:"notification_type,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	CreateTime *int64 `json:"create_time,omitempty"`

	Filter *Filter `json:"filter,omitempty"`
}

func (o NotificationsResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationsResponseBody struct{}"
	}

	return strings.Join([]string{"NotificationsResponseBody", string(data)}, " ")
}

type NotificationsResponseBodyOperationType struct {
	value string
}

type NotificationsResponseBodyOperationTypeEnum struct {
	CUSTOMIZED NotificationsResponseBodyOperationType
	COMPLETE   NotificationsResponseBodyOperationType
}

func GetNotificationsResponseBodyOperationTypeEnum() NotificationsResponseBodyOperationTypeEnum {
	return NotificationsResponseBodyOperationTypeEnum{
		CUSTOMIZED: NotificationsResponseBodyOperationType{
			value: "customized",
		},
		COMPLETE: NotificationsResponseBodyOperationType{
			value: "complete",
		},
	}
}

func (c NotificationsResponseBodyOperationType) Value() string {
	return c.value
}

func (c NotificationsResponseBodyOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NotificationsResponseBodyOperationType) UnmarshalJSON(b []byte) error {
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

type NotificationsResponseBodyStatus struct {
	value string
}

type NotificationsResponseBodyStatusEnum struct {
	ENABLED  NotificationsResponseBodyStatus
	DISABLED NotificationsResponseBodyStatus
}

func GetNotificationsResponseBodyStatusEnum() NotificationsResponseBodyStatusEnum {
	return NotificationsResponseBodyStatusEnum{
		ENABLED: NotificationsResponseBodyStatus{
			value: "enabled",
		},
		DISABLED: NotificationsResponseBodyStatus{
			value: "disabled",
		},
	}
}

func (c NotificationsResponseBodyStatus) Value() string {
	return c.value
}

func (c NotificationsResponseBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NotificationsResponseBodyStatus) UnmarshalJSON(b []byte) error {
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

type NotificationsResponseBodyNotificationType struct {
	value string
}

type NotificationsResponseBodyNotificationTypeEnum struct {
	SMN NotificationsResponseBodyNotificationType
	FUN NotificationsResponseBodyNotificationType
}

func GetNotificationsResponseBodyNotificationTypeEnum() NotificationsResponseBodyNotificationTypeEnum {
	return NotificationsResponseBodyNotificationTypeEnum{
		SMN: NotificationsResponseBodyNotificationType{
			value: "smn",
		},
		FUN: NotificationsResponseBodyNotificationType{
			value: "fun",
		},
	}
}

func (c NotificationsResponseBodyNotificationType) Value() string {
	return c.value
}

func (c NotificationsResponseBodyNotificationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NotificationsResponseBodyNotificationType) UnmarshalJSON(b []byte) error {
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
