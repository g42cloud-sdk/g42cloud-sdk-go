package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateSubscriptionRequest struct {
	TopicUrn string `json:"topic_urn"`

	SubscriptionUrn string `json:"subscription_urn"`

	Body *UpdateSubscriptionRequestBody `json:"body,omitempty"`
}

func (o UpdateSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"UpdateSubscriptionRequest", string(data)}, " ")
}
