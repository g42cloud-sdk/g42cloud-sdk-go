package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type AudioInfo struct {
	Codec *string `json:"codec,omitempty"`

	Sample *int32 `json:"sample,omitempty"`

	Channels *int32 `json:"channels,omitempty"`

	Bitrate *int32 `json:"bitrate,omitempty"`

	BitrateBps *int64 `json:"bitrate_bps,omitempty"`
}

func (o AudioInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AudioInfo struct{}"
	}

	return strings.Join([]string{"AudioInfo", string(data)}, " ")
}
