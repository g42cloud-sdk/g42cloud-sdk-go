package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type Notification struct {
	EventName *string `json:"event_name,omitempty"`

	Status *NotificationStatus `json:"status,omitempty"`

	Topic *string `json:"topic,omitempty"`

	MsgType *int32 `json:"msg_type,omitempty"`
}

func (o Notification) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Notification struct{}"
	}

	return strings.Join([]string{"Notification", string(data)}, " ")
}

type NotificationStatus struct {
	value string
}

type NotificationStatusEnum struct {
	ON  NotificationStatus
	OFF NotificationStatus
}

func GetNotificationStatusEnum() NotificationStatusEnum {
	return NotificationStatusEnum{
		ON: NotificationStatus{
			value: "on",
		},
		OFF: NotificationStatus{
			value: "off",
		},
	}
}

func (c NotificationStatus) Value() string {
	return c.value
}

func (c NotificationStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NotificationStatus) UnmarshalJSON(b []byte) error {
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
