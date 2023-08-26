package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type OriginPara struct {
	Duration *int32 `json:"duration,omitempty"`

	DurationMs *int64 `json:"duration_ms,omitempty"`

	FileFormat *string `json:"file_format,omitempty"`

	Video *VideoInfo `json:"video,omitempty"`

	Audio *AudioInfo `json:"audio,omitempty"`
}

func (o OriginPara) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OriginPara struct{}"
	}

	return strings.Join([]string{"OriginPara", string(data)}, " ")
}
