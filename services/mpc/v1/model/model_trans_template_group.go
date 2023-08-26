package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type TransTemplateGroup struct {
	Name string `json:"name"`

	Videos *[]VideoObj `json:"videos,omitempty"`

	Audio *Audio `json:"audio,omitempty"`

	VideoCommon *VideoCommon `json:"video_common,omitempty"`

	Common *Common `json:"common,omitempty"`
}

func (o TransTemplateGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransTemplateGroup struct{}"
	}

	return strings.Join([]string{"TransTemplateGroup", string(data)}, " ")
}
