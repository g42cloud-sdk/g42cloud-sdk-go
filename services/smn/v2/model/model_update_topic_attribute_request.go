package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateTopicAttributeRequest struct {
	TopicUrn string `json:"topic_urn"`

	Name string `json:"name"`

	Body *UpdateTopicAttributeRequestBody `json:"body,omitempty"`
}

func (o UpdateTopicAttributeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicAttributeRequest struct{}"
	}

	return strings.Join([]string{"UpdateTopicAttributeRequest", string(data)}, " ")
}
