package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceRequirements struct {
	Limits map[string]string `json:"limits,omitempty"`

	Requests map[string]string `json:"requests,omitempty"`
}

func (o ResourceRequirements) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceRequirements struct{}"
	}

	return strings.Join([]string{"ResourceRequirements", string(data)}, " ")
}
