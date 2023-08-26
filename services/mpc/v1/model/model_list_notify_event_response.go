package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListNotifyEventResponse struct {
	EventName *[]string `json:"event_name,omitempty"`

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListNotifyEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotifyEventResponse struct{}"
	}

	return strings.Join([]string{"ListNotifyEventResponse", string(data)}, " ")
}
