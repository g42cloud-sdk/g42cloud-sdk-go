package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
