package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateEditingJobReq struct {
	EditType *[]string `json:"edit_type,omitempty"`

	Clips *[]ClipInfo `json:"clips,omitempty"`

	Concats *[]MultiConcatInfo `json:"concats,omitempty"`

	Concat *ConcatInfo `json:"concat,omitempty"`

	Mix *MixInfo `json:"mix,omitempty"`

	Input *ObsObjInfo `json:"input,omitempty"`

	OutputSetting *OutputSetting `json:"output_setting,omitempty"`

	ImageWatermarkSettings *[]ImageWatermarkSetting `json:"image_watermark_settings,omitempty"`

	EditSettings *[]EditSetting `json:"edit_settings,omitempty"`

	UserData *string `json:"user_data,omitempty"`
}

func (o CreateEditingJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEditingJobReq struct{}"
	}

	return strings.Join([]string{"CreateEditingJobReq", string(data)}, " ")
}
