package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type SubAudioFile struct {
	TracksInfo *[]TracksInfo `json:"tracks_info,omitempty"`

	Input *ObsObjInfo `json:"input,omitempty"`

	Output *ObsObjInfo `json:"output,omitempty"`

	OutputFilename *string `json:"output_filename,omitempty"`
}

func (o SubAudioFile) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubAudioFile struct{}"
	}

	return strings.Join([]string{"SubAudioFile", string(data)}, " ")
}
