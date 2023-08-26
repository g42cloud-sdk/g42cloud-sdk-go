package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ConfigBody struct {
	ConfigKey string `json:"config_key"`

	ConfigValue string `json:"config_value"`

	ConfigStatus *string `json:"config_status,omitempty"`
}

func (o ConfigBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigBody struct{}"
	}

	return strings.Join([]string{"ConfigBody", string(data)}, " ")
}
