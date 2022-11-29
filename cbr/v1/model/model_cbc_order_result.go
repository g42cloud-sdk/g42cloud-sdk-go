package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
