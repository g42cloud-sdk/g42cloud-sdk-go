package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ConfigurationRequestBody struct {
	Configurations []ConfigBody `json:"configurations"`
}

func (o ConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"ConfigurationRequestBody", string(data)}, " ")
}
