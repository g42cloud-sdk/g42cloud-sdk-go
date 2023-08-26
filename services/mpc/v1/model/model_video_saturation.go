package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type VideoSaturation struct {
	Name *string `json:"name,omitempty"`

	ExecutionOrder *int32 `json:"execution_order,omitempty"`

	Saturation *string `json:"saturation,omitempty"`
}

func (o VideoSaturation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoSaturation struct{}"
	}

	return strings.Join([]string{"VideoSaturation", string(data)}, " ")
}
