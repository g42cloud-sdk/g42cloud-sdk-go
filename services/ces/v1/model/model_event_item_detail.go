package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type EventItemDetail struct {
	Content *string `json:"content,omitempty"`

	GroupId *string `json:"group_id,omitempty"`

	ResourceId *string `json:"resource_id,omitempty"`

	ResourceName *string `json:"resource_name,omitempty"`

	EventState *EventItemDetailEventState `json:"event_state,omitempty"`

	EventLevel *EventItemDetailEventLevel `json:"event_level,omitempty"`

	EventUser *string `json:"event_user,omitempty"`

	EventType *string `json:"event_type,omitempty"`
}

func (o EventItemDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventItemDetail struct{}"
	}

	return strings.Join([]string{"EventItemDetail", string(data)}, " ")
}

type EventItemDetailEventState struct {
	value string
}

type EventItemDetailEventStateEnum struct {
	NORMAL   EventItemDetailEventState
	WARNING  EventItemDetailEventState
	INCIDENT EventItemDetailEventState
}

func GetEventItemDetailEventStateEnum() EventItemDetailEventStateEnum {
	return EventItemDetailEventStateEnum{
		NORMAL: EventItemDetailEventState{
			value: "normal",
		},
		WARNING: EventItemDetailEventState{
			value: "warning",
		},
		INCIDENT: EventItemDetailEventState{
			value: "incident",
		},
	}
}

func (c EventItemDetailEventState) Value() string {
	return c.value
}

func (c EventItemDetailEventState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventItemDetailEventState) UnmarshalJSON(b []byte) error {
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

type EventItemDetailEventLevel struct {
	value string
}

type EventItemDetailEventLevelEnum struct {
	CRITICAL EventItemDetailEventLevel
	MAJOR    EventItemDetailEventLevel
	MINOR    EventItemDetailEventLevel
	INFO     EventItemDetailEventLevel
}

func GetEventItemDetailEventLevelEnum() EventItemDetailEventLevelEnum {
	return EventItemDetailEventLevelEnum{
		CRITICAL: EventItemDetailEventLevel{
			value: "Critical",
		},
		MAJOR: EventItemDetailEventLevel{
			value: "Major",
		},
		MINOR: EventItemDetailEventLevel{
			value: "Minor",
		},
		INFO: EventItemDetailEventLevel{
			value: "Info",
		},
	}
}

func (c EventItemDetailEventLevel) Value() string {
	return c.value
}

func (c EventItemDetailEventLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventItemDetailEventLevel) UnmarshalJSON(b []byte) error {
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
