package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BasicWatermark struct {
	Dx *string `json:"dx,omitempty"`

	Dy *string `json:"dy,omitempty"`

	Referpos *string `json:"referpos,omitempty"`

	TimelineStart *string `json:"timeline_start,omitempty"`

	TimelineDuration *string `json:"timeline_duration,omitempty"`
}

func (o BasicWatermark) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasicWatermark struct{}"
	}

	return strings.Join([]string{"BasicWatermark", string(data)}, " ")
}
