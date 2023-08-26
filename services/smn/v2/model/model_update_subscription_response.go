package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateSubscriptionResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	SubscriptionUrn *string `json:"subscription_urn,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o UpdateSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"UpdateSubscriptionResponse", string(data)}, " ")
}
