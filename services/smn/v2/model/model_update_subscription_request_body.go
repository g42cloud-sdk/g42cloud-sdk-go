package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateSubscriptionRequestBody struct {
	Remark string `json:"remark"`
}

func (o UpdateSubscriptionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubscriptionRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSubscriptionRequestBody", string(data)}, " ")
}
