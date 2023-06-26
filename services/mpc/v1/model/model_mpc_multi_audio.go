package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type MpcMultiAudio struct {
	Output *ObsObjInfo `json:"output,omitempty"`

	AudioFiles *[]AudioFile `json:"audio_files,omitempty"`

	OutputFilename *string `json:"output_filename,omitempty"`
}

func (o MpcMultiAudio) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MpcMultiAudio struct{}"
	}

	return strings.Join([]string{"MpcMultiAudio", string(data)}, " ")
}
