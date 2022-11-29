package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EventInfoDetail struct {
	EventName string `json:"event_name"`

	EventSource string `json:"event_source"`

	Time int64 `json:"time"`

	Detail *EventItemDetail `json:"detail"`

	EventId *string `json:"event_id,omitempty"`
}

func (o EventInfoDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventInfoDetail struct{}"
	}

	return strings.Join([]string{"EventInfoDetail", string(data)}, " ")
}
