package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type TemplateGroup struct {
	GroupId *string `json:"group_id,omitempty"`

	Name *string `json:"name,omitempty"`

	TemplateIds *[]int32 `json:"template_ids,omitempty"`

	Videos *[]VideoAndTemplate `json:"videos,omitempty"`

	Audio *Audio `json:"audio,omitempty"`

	VideoCommon *VideoCommon `json:"video_common,omitempty"`

	Common *Common `json:"common,omitempty"`
}

func (o TemplateGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateGroup struct{}"
	}

	return strings.Join([]string{"TemplateGroup", string(data)}, " ")
}
