package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CbcOrderResult struct {
	CloudServiceId *string `json:"cloudServiceId,omitempty"`

	OrderId string `json:"orderId"`

	SubscribeResult int32 `json:"subscribeResult"`

	ResourceId *string `json:"resourceId,omitempty"`
}

func (o CbcOrderResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CbcOrderResult struct{}"
	}

	return strings.Join([]string{"CbcOrderResult", string(data)}, " ")
}
