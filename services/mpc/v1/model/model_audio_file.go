package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type AudioFile struct {
	TracksInfo *[]TracksInfo `json:"tracks_info,omitempty"`

	Input *ObsObjInfo `json:"input,omitempty"`
}

func (o AudioFile) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AudioFile struct{}"
	}

	return strings.Join([]string{"AudioFile", string(data)}, " ")
}
