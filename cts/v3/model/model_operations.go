package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Operations struct {
	ServiceType string `json:"service_type"`

	ResourceType string `json:"resource_type"`

	TraceNames []string `json:"trace_names"`
}

func (o Operations) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Operations struct{}"
	}

	return strings.Join([]string{"Operations", string(data)}, " ")
}
