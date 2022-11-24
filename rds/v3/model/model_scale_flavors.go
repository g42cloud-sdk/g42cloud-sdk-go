package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScaleFlavors struct {
	Code *string `json:"code,omitempty"`

	Cpu *string `json:"cpu,omitempty"`

	Mem *string `json:"mem,omitempty"`
}

func (o ScaleFlavors) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScaleFlavors struct{}"
	}

	return strings.Join([]string{"ScaleFlavors", string(data)}, " ")
}
