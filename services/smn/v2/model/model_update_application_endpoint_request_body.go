package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateApplicationEndpointRequestBody struct {
	Enabled *string `json:"enabled,omitempty"`

	UserData *string `json:"user_data,omitempty"`
}

func (o UpdateApplicationEndpointRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationEndpointRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateApplicationEndpointRequestBody", string(data)}, " ")
}
