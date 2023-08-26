package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteTopicRequest struct {
	TopicUrn string `json:"topic_urn"`
}

func (o DeleteTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTopicRequest struct{}"
	}

	return strings.Join([]string{"DeleteTopicRequest", string(data)}, " ")
}
