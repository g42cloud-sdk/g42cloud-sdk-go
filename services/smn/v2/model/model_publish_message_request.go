package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PublishMessageRequest struct {
	TopicUrn string `json:"topic_urn"`

	Body *PublishMessageRequestBody `json:"body,omitempty"`
}

func (o PublishMessageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishMessageRequest struct{}"
	}

	return strings.Join([]string{"PublishMessageRequest", string(data)}, " ")
}
