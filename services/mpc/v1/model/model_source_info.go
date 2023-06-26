package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type SourceInfo struct {
	Duration *int32 `json:"duration,omitempty"`

	DurationMs *int64 `json:"duration_ms,omitempty"`

	Format *string `json:"format,omitempty"`

	Size *int64 `json:"size,omitempty"`

	VideoInfo *VideoInfo `json:"video_info,omitempty"`

	AudioInfo *[]AudioInfo `json:"audio_info,omitempty"`
}

func (o SourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceInfo struct{}"
	}

	return strings.Join([]string{"SourceInfo", string(data)}, " ")
}
