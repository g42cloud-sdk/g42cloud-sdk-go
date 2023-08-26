package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type TemplateInfo struct {
	TemplateId *int32 `json:"template_id,omitempty"`

	Template *QueryTransTemplate `json:"template,omitempty"`
}

func (o TemplateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateInfo struct{}"
	}

	return strings.Join([]string{"TemplateInfo", string(data)}, " ")
}
