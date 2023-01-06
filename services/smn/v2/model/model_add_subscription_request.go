package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type AddSubscriptionRequest struct {
	TopicUrn string `json:"topic_urn"`

	Body *AddSubscriptionRequestBody `json:"body,omitempty"`
}

func (o AddSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"AddSubscriptionRequest", string(data)}, " ")
}
