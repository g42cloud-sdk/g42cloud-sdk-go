package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type SubscriptionExtension struct {
	ClientId *string `json:"client_id,omitempty"`

	ClientSecret *string `json:"client_secret,omitempty"`

	Keyword *string `json:"keyword,omitempty"`

	SignSecret *string `json:"sign_secret,omitempty"`
}

func (o SubscriptionExtension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionExtension struct{}"
	}

	return strings.Join([]string{"SubscriptionExtension", string(data)}, " ")
}
