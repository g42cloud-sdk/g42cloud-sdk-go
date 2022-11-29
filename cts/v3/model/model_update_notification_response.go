package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateNotificationResponse struct {
	NotificationName *string `json:"notification_name,omitempty"`

	OperationType *UpdateNotificationResponseOperationType `json:"operation_type,omitempty"`

	Operations *[]Operations `json:"operations,omitempty"`

	NotifyUserList *[]NotificationUsers `json:"notify_user_list,omitempty"`

	Status *UpdateNotificationResponseStatus `json:"status,omitempty"`

	TopicId *string `json:"topic_id,omitempty"`

	NotificationId *string `json:"notification_id,omitempty"`

	NotificationType *UpdateNotificationResponseNotificationType `json:"notification_type,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	CreateTime *int64 `json:"create_time,omitempty"`

	Filter         *Filter `json:"filter,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateNotificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotificationResponse struct{}"
	}

	return strings.Join([]string{"UpdateNotificationResponse", string(data)}, " ")
}

type UpdateNotificationResponseOperationType struct {
	value string
}

type UpdateNotificationResponseOperationTypeEnum struct {
	CUSTOMIZED UpdateNotificationResponseOperationType
	COMPLETE   UpdateNotificationResponseOperationType
}

func GetUpdateNotificationResponseOperationTypeEnum() UpdateNotificationResponseOperationTypeEnum {
	return UpdateNotificationResponseOperationTypeEnum{
		CUSTOMIZED: UpdateNotificationResponseOperationType{
			value: "customized",
		},
		COMPLETE: UpdateNotificationResponseOperationType{
			value: "complete",
		},
	}
}

func (c UpdateNotificationResponseOperationType) Value() string {
	return c.value
}

func (c UpdateNotificationResponseOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateNotificationResponseOperationType) UnmarshalJSON(b []byte) error {
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

type UpdateNotificationResponseStatus struct {
	value string
}

type UpdateNotificationResponseStatusEnum struct {
	ENABLED  UpdateNotificationResponseStatus
	DISABLED UpdateNotificationResponseStatus
}

func GetUpdateNotificationResponseStatusEnum() UpdateNotificationResponseStatusEnum {
	return UpdateNotificationResponseStatusEnum{
		ENABLED: UpdateNotificationResponseStatus{
			value: "enabled",
		},
		DISABLED: UpdateNotificationResponseStatus{
			value: "disabled",
		},
	}
}

func (c UpdateNotificationResponseStatus) Value() string {
	return c.value
}

func (c UpdateNotificationResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateNotificationResponseStatus) UnmarshalJSON(b []byte) error {
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

type UpdateNotificationResponseNotificationType struct {
	value string
}

type UpdateNotificationResponseNotificationTypeEnum struct {
	SMN UpdateNotificationResponseNotificationType
	FUN UpdateNotificationResponseNotificationType
}

func GetUpdateNotificationResponseNotificationTypeEnum() UpdateNotificationResponseNotificationTypeEnum {
	return UpdateNotificationResponseNotificationTypeEnum{
		SMN: UpdateNotificationResponseNotificationType{
			value: "smn",
		},
		FUN: UpdateNotificationResponseNotificationType{
			value: "fun",
		},
	}
}

func (c UpdateNotificationResponseNotificationType) Value() string {
	return c.value
}

func (c UpdateNotificationResponseNotificationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateNotificationResponseNotificationType) UnmarshalJSON(b []byte) error {
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
