package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteTopicAttributesRequest struct {
	TopicUrn string `json:"topic_urn"`
}

func (o DeleteTopicAttributesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTopicAttributesRequest struct{}"
	}

	return strings.Join([]string{"DeleteTopicAttributesRequest", string(data)}, " ")
}
