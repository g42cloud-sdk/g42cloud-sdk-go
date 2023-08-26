package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteTopicAttributeByNameRequest struct {
	TopicUrn string `json:"topic_urn"`

	Name string `json:"name"`
}

func (o DeleteTopicAttributeByNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTopicAttributeByNameRequest struct{}"
	}

	return strings.Join([]string{"DeleteTopicAttributeByNameRequest", string(data)}, " ")
}
