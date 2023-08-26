package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ForceRedirectConfig struct {
	Status string `json:"status"`

	Type *string `json:"type,omitempty"`
}

func (o ForceRedirectConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ForceRedirectConfig struct{}"
	}

	return strings.Join([]string{"ForceRedirectConfig", string(data)}, " ")
}
