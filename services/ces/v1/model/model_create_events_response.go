package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateEventsResponse struct {
	Body           *[]CreateEventsResponseBody `json:"body,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o CreateEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEventsResponse struct{}"
	}

	return strings.Join([]string{"CreateEventsResponse", string(data)}, " ")
}
