package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ListEventDetailRequest struct {
	ContentType string `json:"Content-Type"`

	EventName string `json:"event_name"`

	EventType ListEventDetailRequestEventType `json:"event_type"`

	EventSource *string `json:"event_source,omitempty"`

	EventLevel *string `json:"event_level,omitempty"`

	EventUser *string `json:"event_user,omitempty"`

	EventState *string `json:"event_state,omitempty"`

	From *int64 `json:"from,omitempty"`

	To *int64 `json:"to,omitempty"`

	Start *int32 `json:"start,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListEventDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventDetailRequest struct{}"
	}

	return strings.Join([]string{"ListEventDetailRequest", string(data)}, " ")
}

type ListEventDetailRequestEventType struct {
	value string
}

type ListEventDetailRequestEventTypeEnum struct {
	EVENT_SYS    ListEventDetailRequestEventType
	EVENT_CUSTOM ListEventDetailRequestEventType
}

func GetListEventDetailRequestEventTypeEnum() ListEventDetailRequestEventTypeEnum {
	return ListEventDetailRequestEventTypeEnum{
		EVENT_SYS: ListEventDetailRequestEventType{
			value: "EVENT.SYS",
		},
		EVENT_CUSTOM: ListEventDetailRequestEventType{
			value: "EVENT.CUSTOM",
		},
	}
}

func (c ListEventDetailRequestEventType) Value() string {
	return c.value
}

func (c ListEventDetailRequestEventType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventDetailRequestEventType) UnmarshalJSON(b []byte) error {
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
