package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ResetPwdResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetPwdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetPwdResponse struct{}"
	}

	return strings.Join([]string{"ResetPwdResponse", string(data)}, " ")
}
