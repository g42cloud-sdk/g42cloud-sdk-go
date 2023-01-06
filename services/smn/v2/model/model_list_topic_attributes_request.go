package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListTopicAttributesRequest struct {
	TopicUrn string `json:"topic_urn"`

	Name string `json:"name"`
}

func (o ListTopicAttributesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopicAttributesRequest struct{}"
	}

	return strings.Join([]string{"ListTopicAttributesRequest", string(data)}, " ")
}
