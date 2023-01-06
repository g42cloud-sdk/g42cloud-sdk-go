package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListSubscriptionsByTopicResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	SubscriptionCount *int32 `json:"subscription_count,omitempty"`

	Subscriptions  *[]ListSubscriptionsItem `json:"subscriptions,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListSubscriptionsByTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionsByTopicResponse struct{}"
	}

	return strings.Join([]string{"ListSubscriptionsByTopicResponse", string(data)}, " ")
}
