package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateResetTracksReq struct {
	AudioFile *SubAudioFile `json:"audio_file,omitempty"`
}

func (o CreateResetTracksReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResetTracksReq struct{}"
	}

	return strings.Join([]string{"CreateResetTracksReq", string(data)}, " ")
}
