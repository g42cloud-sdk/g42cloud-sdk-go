package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ListEventsRequest struct {
	ContentType string `json:"Content-Type"`

	EventType *ListEventsRequestEventType `json:"event_type,omitempty"`

	EventName *string `json:"event_name,omitempty"`

	From *int64 `json:"from,omitempty"`

	To *int64 `json:"to,omitempty"`

	Start *int32 `json:"start,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventsRequest struct{}"
	}

	return strings.Join([]string{"ListEventsRequest", string(data)}, " ")
}

type ListEventsRequestEventType struct {
	value string
}

type ListEventsRequestEventTypeEnum struct {
	EVENT_SYS    ListEventsRequestEventType
	EVENT_CUSTOM ListEventsRequestEventType
}

func GetListEventsRequestEventTypeEnum() ListEventsRequestEventTypeEnum {
	return ListEventsRequestEventTypeEnum{
		EVENT_SYS: ListEventsRequestEventType{
			value: "EVENT.SYS",
		},
		EVENT_CUSTOM: ListEventsRequestEventType{
			value: "EVENT.CUSTOM",
		},
	}
}

func (c ListEventsRequestEventType) Value() string {
	return c.value
}

func (c ListEventsRequestEventType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventsRequestEventType) UnmarshalJSON(b []byte) error {
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
