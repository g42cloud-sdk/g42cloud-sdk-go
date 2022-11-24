package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePrivateipRequestBody struct {
	Privateips []CreatePrivateipOption `json:"privateips"`
}

func (o CreatePrivateipRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateipRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePrivateipRequestBody", string(data)}, " ")
}
