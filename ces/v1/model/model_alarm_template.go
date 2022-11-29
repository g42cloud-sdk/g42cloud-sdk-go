package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmTemplate struct {
	TemplateName *string `json:"template_name,omitempty"`

	TemplateDescription *string `json:"template_description,omitempty"`

	Namespace *string `json:"namespace,omitempty"`

	DimensionName *string `json:"dimension_name,omitempty"`

	TemplateItems *[]TemplateItem `json:"template_items,omitempty"`

	TemplateId *string `json:"template_id,omitempty"`
}

func (o AlarmTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmTemplate struct{}"
	}

	return strings.Join([]string{"AlarmTemplate", string(data)}, " ")
}
