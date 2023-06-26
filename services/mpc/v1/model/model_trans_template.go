package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type TransTemplate struct {
	TemplateName string `json:"template_name"`

	Video *Video `json:"video,omitempty"`

	Audio *Audio `json:"audio,omitempty"`

	Common *Common `json:"common"`
}

func (o TransTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransTemplate struct{}"
	}

	return strings.Join([]string{"TransTemplate", string(data)}, " ")
}
