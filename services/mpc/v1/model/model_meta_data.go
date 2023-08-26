package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type MetaData struct {
	Size *int64 `json:"size,omitempty"`

	DurationMs *float64 `json:"duration_ms,omitempty"`

	Duration *int64 `json:"duration,omitempty"`

	Format *string `json:"format,omitempty"`

	Bitrate *int64 `json:"bitrate,omitempty"`

	Video *[]VideoInfo `json:"video,omitempty"`

	Audio *[]AudioInfo `json:"audio,omitempty"`
}

func (o MetaData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetaData struct{}"
	}

	return strings.Join([]string{"MetaData", string(data)}, " ")
}
