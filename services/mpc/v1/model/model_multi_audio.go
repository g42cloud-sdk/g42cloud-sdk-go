package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type MultiAudio struct {
	TracksInfo *[]TracksInfo `json:"tracks_info,omitempty"`

	AudioFiles *[]AudioFile `json:"audio_files,omitempty"`

	DefaultLanguage *string `json:"default_language,omitempty"`
}

func (o MultiAudio) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiAudio struct{}"
	}

	return strings.Join([]string{"MultiAudio", string(data)}, " ")
}
