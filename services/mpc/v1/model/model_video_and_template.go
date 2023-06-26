package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type VideoAndTemplate struct {
	TemplateId *int32 `json:"template_id,omitempty"`

	Width *int32 `json:"width,omitempty"`

	Height *int32 `json:"height,omitempty"`

	Bitrate *int32 `json:"bitrate,omitempty"`
}

func (o VideoAndTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoAndTemplate struct{}"
	}

	return strings.Join([]string{"VideoAndTemplate", string(data)}, " ")
}
