package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type EditAudioInfo struct {
	Codec *string `json:"codec,omitempty"`

	Bitrate *int32 `json:"bitrate,omitempty"`

	Sample *int32 `json:"sample,omitempty"`

	Channels *string `json:"channels,omitempty"`
}

func (o EditAudioInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EditAudioInfo struct{}"
	}

	return strings.Join([]string{"EditAudioInfo", string(data)}, " ")
}
