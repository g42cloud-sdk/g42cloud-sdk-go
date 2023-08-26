package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateTopicRequest struct {
	TopicUrn string `json:"topic_urn"`

	Body *UpdateTopicRequestBody `json:"body,omitempty"`
}

func (o UpdateTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicRequest struct{}"
	}

	return strings.Join([]string{"UpdateTopicRequest", string(data)}, " ")
}
