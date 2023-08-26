package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateAlarmTemplateRequestBody struct {
	TemplateName string `json:"template_name"`

	TemplateDescription *string `json:"template_description,omitempty"`

	Namespace string `json:"namespace"`

	DimensionName string `json:"dimension_name"`

	TemplateItems []TemplateItem `json:"template_items"`
}

func (o CreateAlarmTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlarmTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAlarmTemplateRequestBody", string(data)}, " ")
}
