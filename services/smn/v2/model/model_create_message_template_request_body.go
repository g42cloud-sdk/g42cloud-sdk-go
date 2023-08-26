package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateMessageTemplateRequestBody struct {
	MessageTemplateName string `json:"message_template_name"`

	Protocol string `json:"protocol"`

	Content string `json:"content"`
}

func (o CreateMessageTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMessageTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateMessageTemplateRequestBody", string(data)}, " ")
}
