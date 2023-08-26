package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CancelSubscriptionRequest struct {
	SubscriptionUrn string `json:"subscription_urn"`
}

func (o CancelSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"CancelSubscriptionRequest", string(data)}, " ")
}
