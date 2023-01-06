package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateEventsRequest struct {
	ContentType string `json:"Content-Type"`

	Body *[]EventItem `json:"body,omitempty"`
}

func (o CreateEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEventsRequest struct{}"
	}

	return strings.Join([]string{"CreateEventsRequest", string(data)}, " ")
}
