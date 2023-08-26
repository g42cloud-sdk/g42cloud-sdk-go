package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PublishMessageResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	MessageId      *string `json:"message_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PublishMessageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishMessageResponse struct{}"
	}

	return strings.Join([]string{"PublishMessageResponse", string(data)}, " ")
}
