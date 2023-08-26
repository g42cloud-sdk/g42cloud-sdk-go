package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type EditSetting struct {
	TimelineStart *string `json:"timeline_start,omitempty"`

	TimelineEnd *string `json:"timeline_end,omitempty"`

	TransTemplateId *int32 `json:"trans_template_id,omitempty"`

	AvParameter *AvParameters `json:"av_parameter,omitempty"`

	Mosaics *[]MosaicInfo `json:"mosaics,omitempty"`

	ImageWatermarks *[]ImageWatermarkSetting `json:"image_watermarks,omitempty"`

	Heads *[]ObsObjInfo `json:"heads,omitempty"`

	Tails *[]ObsObjInfo `json:"tails,omitempty"`

	Output *ObsObjInfo `json:"output"`
}

func (o EditSetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EditSetting struct{}"
	}

	return strings.Join([]string{"EditSetting", string(data)}, " ")
}
