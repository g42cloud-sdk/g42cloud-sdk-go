package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type WatermarkRequest struct {
	Input *ObsObjInfo `json:"input,omitempty"`

	TemplateId *string `json:"template_id,omitempty"`

	TextContext *string `json:"text_context,omitempty"`

	ImageWatermark *ImageWatermark `json:"image_watermark,omitempty"`

	TextWatermark *TextWatermark `json:"text_watermark,omitempty"`
}

func (o WatermarkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WatermarkRequest struct{}"
	}

	return strings.Join([]string{"WatermarkRequest", string(data)}, " ")
}
