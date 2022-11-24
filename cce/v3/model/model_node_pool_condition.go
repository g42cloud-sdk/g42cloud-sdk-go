package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodePoolCondition struct {
	Type *string `json:"type,omitempty"`

	Status *string `json:"status,omitempty"`

	LastProbeTime *string `json:"lastProbeTime,omitempty"`

	LastTransitTime *string `json:"lastTransitTime,omitempty"`

	Reason *string `json:"reason,omitempty"`

	Message *string `json:"message,omitempty"`
}

func (o NodePoolCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodePoolCondition struct{}"
	}

	return strings.Join([]string{"NodePoolCondition", string(data)}, " ")
}
