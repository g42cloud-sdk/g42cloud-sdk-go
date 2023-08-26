package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type MultiConcatInfo struct {
	Inputs []ObsObjInfo `json:"inputs"`

	TransTemplateIds *[]int32 `json:"trans_template_ids,omitempty"`

	AvParameters *[]AvParameters `json:"av_parameters,omitempty"`

	Output *ObsObjInfo `json:"output"`

	ImageWatermarkSettings *[]ImageWatermarkSetting `json:"image_watermark_settings,omitempty"`
}

func (o MultiConcatInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiConcatInfo struct{}"
	}

	return strings.Join([]string{"MultiConcatInfo", string(data)}, " ")
}
