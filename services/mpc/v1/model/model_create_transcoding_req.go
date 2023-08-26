package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateTranscodingReq struct {
	Input *ObsObjInfo `json:"input,omitempty"`

	Output *ObsObjInfo `json:"output"`

	TransTemplateId *[]int32 `json:"trans_template_id,omitempty"`

	AvParameters *[]AvParameters `json:"av_parameters,omitempty"`

	OutputFilenames *[]string `json:"output_filenames,omitempty"`

	UserData *string `json:"user_data,omitempty"`

	Watermarks *[]WatermarkRequest `json:"watermarks,omitempty"`

	Thumbnail *Thumbnail `json:"thumbnail,omitempty"`

	Priority *int32 `json:"priority,omitempty"`

	Subtitle *Subtitle `json:"subtitle,omitempty"`

	Encryption *Encryption `json:"encryption,omitempty"`

	Crop *Crop `json:"crop,omitempty"`

	AudioTrack *AudioTrack `json:"audio_track,omitempty"`

	MultiAudio *MultiAudio `json:"multi_audio,omitempty"`

	VideoProcess *VideoProcess `json:"video_process,omitempty"`

	AudioProcess *AudioProcess `json:"audio_process,omitempty"`
}

func (o CreateTranscodingReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTranscodingReq struct{}"
	}

	return strings.Join([]string{"CreateTranscodingReq", string(data)}, " ")
}
