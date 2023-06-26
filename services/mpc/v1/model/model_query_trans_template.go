package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type QueryTransTemplate struct {
	TemplateName string `json:"template_name"`

	Video *Video `json:"video"`

	Audio *Audio `json:"audio,omitempty"`

	Common *Common `json:"common,omitempty"`
}

func (o QueryTransTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryTransTemplate struct{}"
	}

	return strings.Join([]string{"QueryTransTemplate", string(data)}, " ")
}
