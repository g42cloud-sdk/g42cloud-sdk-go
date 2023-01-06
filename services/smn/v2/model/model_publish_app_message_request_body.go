package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PublishAppMessageRequestBody struct {
	Message *string `json:"message,omitempty"`

	MessageStructure *string `json:"message_structure,omitempty"`

	TimeToLive *string `json:"time_to_live,omitempty"`
}

func (o PublishAppMessageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishAppMessageRequestBody struct{}"
	}

	return strings.Join([]string{"PublishAppMessageRequestBody", string(data)}, " ")
}
