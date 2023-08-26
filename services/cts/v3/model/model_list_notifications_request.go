package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ListNotificationsRequest struct {
	NotificationType ListNotificationsRequestNotificationType `json:"notification_type"`

	NotificationName *string `json:"notification_name,omitempty"`
}

func (o ListNotificationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationsRequest struct{}"
	}

	return strings.Join([]string{"ListNotificationsRequest", string(data)}, " ")
}

type ListNotificationsRequestNotificationType struct {
	value string
}

type ListNotificationsRequestNotificationTypeEnum struct {
	SMN ListNotificationsRequestNotificationType
	FUN ListNotificationsRequestNotificationType
}

func GetListNotificationsRequestNotificationTypeEnum() ListNotificationsRequestNotificationTypeEnum {
	return ListNotificationsRequestNotificationTypeEnum{
		SMN: ListNotificationsRequestNotificationType{
			value: "smn",
		},
		FUN: ListNotificationsRequestNotificationType{
			value: "fun",
		},
	}
}

func (c ListNotificationsRequestNotificationType) Value() string {
	return c.value
}

func (c ListNotificationsRequestNotificationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNotificationsRequestNotificationType) UnmarshalJSON(b []byte) error {
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
