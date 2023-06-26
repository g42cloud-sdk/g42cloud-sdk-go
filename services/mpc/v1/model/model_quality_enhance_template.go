package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type QualityEnhanceTemplate struct {
	TemplateName string `json:"template_name"`

	TemplateDescription *string `json:"template_description,omitempty"`

	Video *QualityEnhanceVideo `json:"video,omitempty"`
}

func (o QualityEnhanceTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QualityEnhanceTemplate struct{}"
	}

	return strings.Join([]string{"QualityEnhanceTemplate", string(data)}, " ")
}
