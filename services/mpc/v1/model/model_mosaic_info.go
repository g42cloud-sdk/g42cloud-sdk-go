package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type MosaicInfo struct {
	TimelineStart *string `json:"timeline_start,omitempty"`

	TimelineDuration *string `json:"timeline_duration,omitempty"`

	Dx *string `json:"dx,omitempty"`

	Dy *string `json:"dy,omitempty"`

	Width string `json:"width"`

	Height string `json:"height"`
}

func (o MosaicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MosaicInfo struct{}"
	}

	return strings.Join([]string{"MosaicInfo", string(data)}, " ")
}
