package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateConfigurationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigurationResponse struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationResponse", string(data)}, " ")
}
