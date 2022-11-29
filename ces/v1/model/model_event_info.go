package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EventInfo struct {
	EventName *string `json:"event_name,omitempty"`

	EventType *string `json:"event_type,omitempty"`

	EventCount *int32 `json:"event_count,omitempty"`

	LatestOccurTime *int64 `json:"latest_occur_time,omitempty"`

	LatestEventSource *string `json:"latest_event_source,omitempty"`
}

func (o EventInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventInfo struct{}"
	}

	return strings.Join([]string{"EventInfo", string(data)}, " ")
}
