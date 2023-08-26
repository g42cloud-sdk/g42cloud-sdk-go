package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ModifyTransTemplateReq struct {
	TemplateId int64 `json:"template_id"`

	TemplateName string `json:"template_name"`

	Video *Video `json:"video,omitempty"`

	Audio *Audio `json:"audio,omitempty"`

	Common *Common `json:"common"`
}

func (o ModifyTransTemplateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyTransTemplateReq struct{}"
	}

	return strings.Join([]string{"ModifyTransTemplateReq", string(data)}, " ")
}
