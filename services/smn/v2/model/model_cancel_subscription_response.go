package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CancelSubscriptionResponse struct {
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CancelSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"CancelSubscriptionResponse", string(data)}, " ")
}
