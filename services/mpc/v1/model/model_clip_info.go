package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ClipInfo struct {
	Input *ObsObjInfo `json:"input,omitempty"`

	TimelineStart *string `json:"timeline_start,omitempty"`

	TimelineEnd *string `json:"timeline_end,omitempty"`
}

func (o ClipInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClipInfo struct{}"
	}

	return strings.Join([]string{"ClipInfo", string(data)}, " ")
}
