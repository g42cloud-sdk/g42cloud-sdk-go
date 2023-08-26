package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type AddSubscriptionResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	SubscriptionUrn *string `json:"subscription_urn,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o AddSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"AddSubscriptionResponse", string(data)}, " ")
}
