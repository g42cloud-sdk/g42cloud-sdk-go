package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ModifyTransTemplateGroup struct {
	GroupId string `json:"group_id"`

	Name *string `json:"name,omitempty"`

	Videos *[]VideoObj `json:"videos,omitempty"`

	Audio *Audio `json:"audio,omitempty"`

	VideoCommon *VideoCommon `json:"video_common,omitempty"`

	Common *Common `json:"common,omitempty"`
}

func (o ModifyTransTemplateGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyTransTemplateGroup struct{}"
	}

	return strings.Join([]string{"ModifyTransTemplateGroup", string(data)}, " ")
}
