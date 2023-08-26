package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateApplicationEndpointRequestBody struct {
	Token string `json:"token"`

	UserData string `json:"user_data"`
}

func (o CreateApplicationEndpointRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationEndpointRequestBody struct{}"
	}

	return strings.Join([]string{"CreateApplicationEndpointRequestBody", string(data)}, " ")
}
