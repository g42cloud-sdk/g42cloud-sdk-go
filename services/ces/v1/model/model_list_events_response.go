package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListEventsResponse struct {
	Events *[]EventInfo `json:"events,omitempty"`

	MetaData       *TotalMetaData `json:"meta_data,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventsResponse struct{}"
	}

	return strings.Join([]string{"ListEventsResponse", string(data)}, " ")
}
