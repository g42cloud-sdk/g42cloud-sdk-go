package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type AddSubscriptionRequestBody struct {
	Protocol string `json:"protocol"`

	Endpoint string `json:"endpoint"`

	Remark *string `json:"remark,omitempty"`

	Extension *SubscriptionExtension `json:"extension,omitempty"`
}

func (o AddSubscriptionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSubscriptionRequestBody struct{}"
	}

	return strings.Join([]string{"AddSubscriptionRequestBody", string(data)}, " ")
}
