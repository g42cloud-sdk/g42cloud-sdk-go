package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type MpeMetaData struct {
	PackType *string `json:"pack_type,omitempty"`

	Duration *float64 `json:"duration,omitempty"`

	VideoSize *int64 `json:"video_size,omitempty"`

	Width *int32 `json:"width,omitempty"`

	Height *int32 `json:"height,omitempty"`

	BitRate *int32 `json:"bit_rate,omitempty"`

	AudioBitRate *int32 `json:"audio_bit_rate,omitempty"`

	FrameRate *int32 `json:"frame_rate,omitempty"`

	CodecName *string `json:"codec_name,omitempty"`

	AudioCodecName *string `json:"audio_codec_name,omitempty"`

	Channels *int32 `json:"channels,omitempty"`

	Sample *int32 `json:"sample,omitempty"`

	IsAudio *bool `json:"is_audio,omitempty"`
}

func (o MpeMetaData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MpeMetaData struct{}"
	}

	return strings.Join([]string{"MpeMetaData", string(data)}, " ")
}
