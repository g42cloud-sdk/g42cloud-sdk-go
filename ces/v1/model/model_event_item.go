package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EventItem struct {
	EventName string `json:"event_name"`

	EventSource string `json:"event_source"`

	Time int64 `json:"time"`

	Detail *EventItemDetail `json:"detail"`
}

func (o EventItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventItem struct{}"
	}

	return strings.Join([]string{"EventItem", string(data)}, " ")
}
