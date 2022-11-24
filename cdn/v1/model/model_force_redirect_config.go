package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
