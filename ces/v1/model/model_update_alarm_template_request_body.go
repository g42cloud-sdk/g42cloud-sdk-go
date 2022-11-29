package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAlarmTemplateRequestBody struct {
	TemplateName string `json:"template_name"`

	TemplateDescription *string `json:"template_description,omitempty"`

	Namespace string `json:"namespace"`

	DimensionName string `json:"dimension_name"`

	TemplateItems []TemplateItem `json:"template_items"`
}

func (o UpdateAlarmTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAlarmTemplateRequestBody", string(data)}, " ")
}
