package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListTopicDetailsRequest struct {
	TopicUrn string `json:"topic_urn"`
}

func (o ListTopicDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopicDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListTopicDetailsRequest", string(data)}, " ")
}
