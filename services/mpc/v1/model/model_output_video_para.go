package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type OutputVideoPara struct {
	TemplateId *int32 `json:"template_id,omitempty"`

	Size *int64 `json:"size,omitempty"`

	Pack *string `json:"pack,omitempty"`

	Video *VideoInfo `json:"video,omitempty"`

	Audio *AudioInfo `json:"audio,omitempty"`

	FileName *string `json:"file_name,omitempty"`

	ConverDuration *float64 `json:"conver_duration,omitempty"`

	Error *XCodeError `json:"error,omitempty"`
}

func (o OutputVideoPara) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OutputVideoPara struct{}"
	}

	return strings.Join([]string{"OutputVideoPara", string(data)}, " ")
}
