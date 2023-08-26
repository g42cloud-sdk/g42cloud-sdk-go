package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PublishAppMessageResponse struct {
	MessageId *string `json:"message_id,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PublishAppMessageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishAppMessageResponse struct{}"
	}

	return strings.Join([]string{"PublishAppMessageResponse", string(data)}, " ")
}
