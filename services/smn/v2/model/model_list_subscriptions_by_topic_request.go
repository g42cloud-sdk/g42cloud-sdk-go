package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListSubscriptionsByTopicRequest struct {
	TopicUrn string `json:"topic_urn"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSubscriptionsByTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionsByTopicRequest struct{}"
	}

	return strings.Join([]string{"ListSubscriptionsByTopicRequest", string(data)}, " ")
}
