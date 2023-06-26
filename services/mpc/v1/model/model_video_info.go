package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type VideoInfo struct {
	Width *int32 `json:"width,omitempty"`

	Height *int32 `json:"height,omitempty"`

	Bitrate *int32 `json:"bitrate,omitempty"`

	BitrateBps *int64 `json:"bitrate_bps,omitempty"`

	FrameRate *int32 `json:"frame_rate,omitempty"`

	Codec *string `json:"codec,omitempty"`
}

func (o VideoInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoInfo struct{}"
	}

	return strings.Join([]string{"VideoInfo", string(data)}, " ")
}
