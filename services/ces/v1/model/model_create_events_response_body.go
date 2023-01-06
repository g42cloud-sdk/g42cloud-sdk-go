package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateEventsResponseBody struct {
	EventId string `json:"event_id"`

	EventName string `json:"event_name"`
}

func (o CreateEventsResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEventsResponseBody struct{}"
	}

	return strings.Join([]string{"CreateEventsResponseBody", string(data)}, " ")
}
