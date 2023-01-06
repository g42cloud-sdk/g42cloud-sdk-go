package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListSubscriptionsItem struct {
	TopicUrn string `json:"topic_urn"`

	Protocol string `json:"protocol"`

	SubscriptionUrn string `json:"subscription_urn"`

	Owner string `json:"owner"`

	Endpoint string `json:"endpoint"`

	Remark string `json:"remark"`

	Status int32 `json:"status"`
}

func (o ListSubscriptionsItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionsItem struct{}"
	}

	return strings.Join([]string{"ListSubscriptionsItem", string(data)}, " ")
}
